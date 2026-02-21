package theme

import (
	"math"
	"math/rand"
)

// RGBTo256 converts RGB (0-255) to the nearest 256-color index.
// Ported from play.js:34-46.
func RGBTo256(r, g, b int) int {
	if r == g && g == b {
		if r < 8 {
			return 16
		}
		if r > 248 {
			return 231
		}
		return int(math.Round(float64(r-8)/247*24)) + 232
	}
	r6 := int(math.Round(float64(r) / 255 * 5))
	g6 := int(math.Round(float64(g) / 255 * 5))
	b6 := int(math.Round(float64(b) / 255 * 5))
	return 16 + r6*36 + g6*6 + b6
}

// HSL represents a color in HSL space.
type HSL struct {
	H float64 // 0-360
	S float64 // 0-100
	L float64 // 0-100
}

// HSLToRGB converts HSL to RGB (0-255).
func HSLToRGB(h, s, l float64) (int, int, int) {
	s /= 100
	l /= 100

	c := (1 - math.Abs(2*l-1)) * s
	x := c * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := l - c/2

	var r1, g1, b1 float64
	switch {
	case h < 60:
		r1, g1, b1 = c, x, 0
	case h < 120:
		r1, g1, b1 = x, c, 0
	case h < 180:
		r1, g1, b1 = 0, c, x
	case h < 240:
		r1, g1, b1 = 0, x, c
	case h < 300:
		r1, g1, b1 = x, 0, c
	default:
		r1, g1, b1 = c, 0, x
	}

	return int(math.Round((r1 + m) * 255)),
		int(math.Round((g1 + m) * 255)),
		int(math.Round((b1 + m) * 255))
}

// HSLTo256 converts HSL to a 256-color index.
func HSLTo256(h, s, l float64) int {
	r, g, b := HSLToRGB(h, s, l)
	return RGBTo256(r, g, b)
}

// GenerateGradient creates an HSL gradient between two colors.
// Uses shortest-path hue interpolation. Ported from app.js:59-83.
func GenerateGradient(bottom, top HSL, steps int) []int {
	if steps < 2 {
		steps = 2
	}

	colors := make([]int, steps)
	for i := 0; i < steps; i++ {
		t := float64(i) / float64(steps-1)

		hDiff := top.H - bottom.H
		if hDiff > 180 {
			hDiff -= 360
		}
		if hDiff < -180 {
			hDiff += 360
		}
		h := math.Mod(bottom.H+hDiff*t+360, 360)
		s := bottom.S + (top.S-bottom.S)*t
		l := bottom.L + (top.L-bottom.L)*t

		colors[i] = HSLTo256(h, s, l)
	}
	return colors
}

// RandomHSLPair generates a random HSL pair suitable for equalizer gradients.
func RandomHSLPair() (HSL, HSL) {
	h1 := rand.Float64() * 360
	h2 := math.Mod(h1+60+rand.Float64()*180, 360)
	return HSL{H: h1, S: 60 + rand.Float64()*30, L: 50 + rand.Float64()*20},
		HSL{H: h2, S: 60 + rand.Float64()*30, L: 50 + rand.Float64()*20}
}
