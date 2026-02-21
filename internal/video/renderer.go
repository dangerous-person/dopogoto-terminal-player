package video

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dangerous-person/dopogoto/internal/theme"
)

// Render modes.
const (
	RenderNormal    = 0
	RenderGrayscale = 1
	RenderTint      = 2
)

// Renderer converts a decoded video buffer to an ANSI string.
// Ported from play.js:114-129.
type Renderer struct {
	palette    []string // original hex palette (for tinting)
	ansiColors []string // pre-computed ANSI escape per palette entry
	grayColors []string // grayscale version of each palette entry
	tintColors []string // tinted version (amber, etc.)
	tintHue    float64  // cached hue
	tintSat    float64  // cached saturation
}

// NewRenderer creates a renderer from a palette of hex color strings.
func NewRenderer(palette []string) *Renderer {
	colors := make([]string, len(palette))
	grays := make([]string, len(palette))
	for i, hex := range palette {
		r, g, b := parseHex(hex)
		colors[i] = fmt.Sprintf("\x1b[38;5;%dm", theme.RGBTo256(r, g, b))
		// Luminance → 232-255 grayscale ramp (24 shades)
		lum := int(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))
		grays[i] = fmt.Sprintf("\x1b[38;5;%dm", theme.RGBTo256(lum, lum, lum))
	}
	return &Renderer{palette: palette, ansiColors: colors, grayColors: grays}
}

// SetTint builds a tinted palette mapping luminance to a single hue.
// Hue 0-360, sat 0-100. Cached — only rebuilds when hue/sat change.
func (re *Renderer) SetTint(hue, sat float64) {
	if hue == re.tintHue && sat == re.tintSat && len(re.tintColors) == len(re.palette) {
		return
	}
	re.tintHue = hue
	re.tintSat = sat
	re.tintColors = make([]string, len(re.palette))
	for i, hex := range re.palette {
		r, g, b := parseHex(hex)
		lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
		// Map luminance (0-255) to lightness (0-55%) for phosphor look
		l := lum / 255.0 * 55.0
		c256 := theme.HSLTo256(hue, sat, l)
		re.tintColors[i] = fmt.Sprintf("\x1b[38;5;%dm", c256)
	}
}

// Render converts the decoder's buffer to an ANSI string.
// maxW and maxH clamp the output to fit the available panel size.
// mode: RenderNormal, RenderGrayscale, or RenderTint.
func (r *Renderer) Render(d *Decoder, maxW, maxH int, mode int) string {
	w := d.Width()
	h := d.Height()

	renderW := w
	if maxW > 0 && renderW > maxW {
		renderW = maxW
	}
	renderH := h
	if maxH > 0 && renderH > maxH {
		renderH = maxH
	}

	var b strings.Builder
	b.Grow(renderW * renderH * 12) // rough estimate

	palette := r.ansiColors
	switch mode {
	case RenderGrayscale:
		palette = r.grayColors
	case RenderTint:
		palette = r.tintColors
	}

	lastColorIdx := -1
	for y := 0; y < renderH; y++ {
		for x := 0; x < renderW; x++ {
			cell := d.Buffer[y*w+x]
			if cell.ColorIdx != lastColorIdx {
				if cell.ColorIdx >= 0 && cell.ColorIdx < len(palette) {
					b.WriteString(palette[cell.ColorIdx])
				}
				lastColorIdx = cell.ColorIdx
			}
			b.WriteRune(d.Char(cell.CharIdx))
		}
		if y < renderH-1 {
			b.WriteByte('\n')
		}
	}
	b.WriteString("\x1b[0m")
	return b.String()
}

// parseHex parses a hex color string like "#ff00aa" to RGB.
func parseHex(hex string) (int, int, int) {
	if len(hex) > 0 && hex[0] == '#' {
		hex = hex[1:]
	}
	if len(hex) != 6 {
		return 0, 0, 0
	}
	r, _ := strconv.ParseInt(hex[0:2], 16, 32)
	g, _ := strconv.ParseInt(hex[2:4], 16, 32)
	b, _ := strconv.ParseInt(hex[4:6], 16, 32)
	return int(r), int(g), int(b)
}
