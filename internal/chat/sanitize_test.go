package chat

import "testing"

func TestSanitize(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"plain ascii", "hello world", "hello world"},
		{"empty", "", ""},
		{"emoji removed", "hello 🌍 world", "hello  world"},
		{"variation selector stripped", "text\uFE0Fmore", "textmore"},
		{"ZWJ stripped", "a\u200Db", "ab"},
		{"narrow unicode kept", "café résumé", "café résumé"},
		{"cyrillic kept", "Привет", "Привет"},
		{"CJK removed", "hello 你好", "hello "},
		{"mixed", "user🎵: hey! 👋", "user: hey! "},

		// ANSI escape injection attacks — \x1b is stripped, remaining brackets are harmless
		{"ANSI escape stripped", "hello\x1b[2Jworld", "hello[2Jworld"},
		{"ANSI color stripped", "\x1b[38;5;196mred text\x1b[0m", "[38;5;196mred text[0m"},
		{"terminal title attack", "\x1b]0;pwned\x07title", "]0;pwnedtitle"},
		{"bell char stripped", "ding\x07dong", "dingdong"},
		{"backspace stripped", "abc\x08\x08XY", "abcXY"},
		{"null byte stripped", "hel\x00lo", "hello"},
		{"tab stripped", "col1\tcol2", "col1col2"},
		{"newline stripped", "line1\nline2", "line1line2"},
		{"DEL stripped", "test\x7Fmore", "testmore"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Sanitize(tt.input)
			if got != tt.want {
				t.Errorf("Sanitize(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}
