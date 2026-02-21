package panels

import (
	"fmt"
	"strings"
)

const FadeDashes = 2 // how many ─/│ near corners use fade color

// FadeBorder renders n dashes where the first `fade` and last `fade` use fadeCol,
// the rest use borderCol.
func FadeBorder(n int, fade int, fadeCol, borderCol string) string {
	if n <= 0 {
		return ""
	}
	if n <= fade*2 {
		return fmt.Sprintf("\x1b[38;5;%sm%s", fadeCol, strings.Repeat("─", n))
	}
	return fmt.Sprintf("\x1b[38;5;%sm%s\x1b[38;5;%sm%s\x1b[38;5;%sm%s",
		fadeCol, strings.Repeat("─", fade),
		borderCol, strings.Repeat("─", n-fade*2),
		fadeCol, strings.Repeat("─", fade))
}

// BuildTitleGradient applies the standard title gradient: 1st char c1, 2nd c2, rest c3.
func BuildTitleGradient(s, c1, c2, c3 string) string {
	runes := []rune(s)
	if len(runes) >= 2 {
		return fmt.Sprintf("\x1b[38;5;%sm%c\x1b[38;5;%sm%c\x1b[38;5;%sm%s\x1b[0m",
			c1, runes[0], c2, runes[1], c3, string(runes[2:]))
	}
	if len(runes) == 1 {
		return fmt.Sprintf("\x1b[38;5;%sm%c\x1b[0m", c1, runes[0])
	}
	return ""
}

// AnsiVisLen returns the visible length of a string containing ANSI escapes.
func AnsiVisLen(s string) int {
	n := 0
	inEscape := false
	for _, r := range s {
		if r == '\x1b' {
			inEscape = true
			continue
		}
		if inEscape {
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
				inEscape = false
			}
			continue
		}
		n++
	}
	return n
}

// truncate shortens a string to maxLen runes, appending "..." if truncated.
func truncate(s string, maxLen int) string {
	if maxLen <= 0 {
		return ""
	}
	runes := []rune(s)
	if len(runes) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return string(runes[:maxLen])
	}
	return string(runes[:maxLen-3]) + "..."
}

// writeBorderedLine writes a content line between │ borders with fade logic.
// If padded is true, adds 1 space padding on each side of content.
func writeBorderedLine(b *strings.Builder, borderColor, fadeColor, content string, contentW, row, totalRows int, padded bool) {
	sideColor := borderColor
	if row < FadeDashes || row >= totalRows-FadeDashes {
		sideColor = fadeColor
	}
	visLen := AnsiVisLen(content)
	pad := contentW - visLen
	if pad < 0 {
		pad = 0
	}
	if padded {
		b.WriteString(fmt.Sprintf("\x1b[38;5;%sm│\x1b[0m %s%s \x1b[38;5;%sm│\x1b[0m\n",
			sideColor, content, strings.Repeat(" ", pad), sideColor))
	} else {
		b.WriteString(fmt.Sprintf("\x1b[38;5;%sm│\x1b[0m%s%s\x1b[38;5;%sm│\x1b[0m\n",
			sideColor, content, strings.Repeat(" ", pad), sideColor))
	}
}
