package chat

import (
	"strings"
	"unicode"
)

// Sanitize removes ANSI escape sequences, control characters, emoji, and other
// wide characters from a string — keeping only terminal-safe narrow printables.
func Sanitize(s string) string {
	// Fast path: check if any rune needs stripping
	for _, r := range s {
		if r < 0x20 || r == 0x7F || r > 0x7E {
			return stripSlow(s)
		}
	}
	return s
}

func stripSlow(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	for _, r := range s {
		// Skip variation selectors, ZWJ, and other invisible modifiers
		if r == 0xFE0F || r == 0xFE0E || r == 0x200D || r == 0x20E3 {
			continue
		}
		// Keep printable ASCII (0x20 space through 0x7E tilde)
		// Drop control chars: 0x00-0x1F (includes \x1b ESC) and 0x7F DEL
		if r >= 0x20 && r <= 0x7E {
			b.WriteRune(r)
			continue
		}
		if r < 0x20 || r == 0x7F {
			continue
		}
		// Keep narrow Unicode (Latin extended, Cyrillic, Greek, etc.)
		if isNarrowUnicode(r) && unicode.IsPrint(r) {
			b.WriteRune(r)
			continue
		}
		// Everything else (emojis, CJK, symbols) — drop
	}
	return b.String()
}

// isNarrowUnicode returns true for non-ASCII characters that are
// single-cell width in terminals.
func isNarrowUnicode(r rune) bool {
	return r < 0x2600 // Covers Latin, Cyrillic, Greek, Armenian, Hebrew, Arabic, etc.
}
