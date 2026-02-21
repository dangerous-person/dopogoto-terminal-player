package panels

import (
	"strings"
	"testing"
)

func TestAnsiVisLen(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"empty", "", 0},
		{"plain ascii", "hello", 5},
		{"single escape", "\x1b[38;5;231mhello\x1b[0m", 5},
		{"multiple escapes", "\x1b[38;5;1mA\x1b[38;5;2mB\x1b[0m", 2},
		{"escape only", "\x1b[0m", 0},
		{"mixed content", "  \x1b[48;5;16m\x1b[38;5;231mAlbum Title\x1b[0m  ", 15},
		{"unicode runes", "café", 4},
		{"border chars", "╭──╮", 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AnsiVisLen(tt.input)
			if got != tt.want {
				t.Errorf("AnsiVisLen(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		maxLen int
		want   string
	}{
		{"short enough", "hello", 10, "hello"},
		{"exact length", "hello", 5, "hello"},
		{"needs truncation", "A Song To Enjoy Your Insomnia", 15, "A Song To En..."},
		{"very short max", "hello", 3, "hel"},
		{"max 0", "hello", 0, ""},
		{"empty string", "", 5, ""},
		{"unicode", "café latte", 6, "caf..."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := truncate(tt.input, tt.maxLen)
			if got != tt.want {
				t.Errorf("truncate(%q, %d) = %q, want %q", tt.input, tt.maxLen, got, tt.want)
			}
		})
	}
}

func TestFadeBorder(t *testing.T) {
	tests := []struct {
		name string
		n    int
		fade int
	}{
		{"zero width", 0, 2},
		{"narrow (less than 2*fade)", 3, 2},
		{"normal width", 10, 2},
		{"wide", 30, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FadeBorder(tt.n, tt.fade, "240", "237")
			// Count visible dashes
			visLen := AnsiVisLen(got)
			if visLen != tt.n {
				t.Errorf("FadeBorder(%d, %d) visible length = %d, want %d", tt.n, tt.fade, visLen, tt.n)
			}
		})
	}
}

func TestBuildTitleGradient(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		wantVis string
	}{
		{"normal title", "Albums", "Albums"},
		{"single char", "A", "A"},
		{"empty", "", ""},
		{"two chars", "OK", "OK"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildTitleGradient(tt.text, "231", "229", "226")
			visLen := AnsiVisLen(got)
			wantLen := len([]rune(tt.wantVis))
			if visLen != wantLen {
				t.Errorf("BuildTitleGradient(%q) visible length = %d, want %d", tt.text, visLen, wantLen)
			}
		})
	}
}

func TestWriteBorderedLine(t *testing.T) {
	// Unpadded variant (albumlist/tracklist style)
	var b strings.Builder
	writeBorderedLine(&b, "240", "237", "hello", 10, 1, 5, false)
	line := b.String()

	if !strings.Contains(line, "│") {
		t.Error("writeBorderedLine should contain border chars")
	}
	if !strings.Contains(line, "hello") {
		t.Error("writeBorderedLine should contain content")
	}
	if !strings.HasSuffix(line, "\n") {
		t.Error("writeBorderedLine should end with newline")
	}

	// Padded variant (controls style)
	b.Reset()
	writeBorderedLine(&b, "240", "237", "test", 10, 0, 1, true)
	padded := b.String()

	if !strings.Contains(padded, "│") {
		t.Error("padded writeBorderedLine should contain border chars")
	}
}
