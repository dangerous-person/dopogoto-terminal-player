package chat

import (
	"strings"
	"testing"
)

func TestGenerateAnonName(t *testing.T) {
	seen := make(map[string]bool)
	for i := 0; i < 50; i++ {
		name := GenerateAnonName()
		if !strings.HasPrefix(name, "cli_anon_") {
			t.Errorf("GenerateAnonName() = %q, want prefix 'cli_anon_'", name)
		}
		seen[name] = true
	}
	// With 2-4 digit random numbers, 50 names should produce at least a few unique
	if len(seen) < 5 {
		t.Errorf("GenerateAnonName() produced only %d unique names in 50 calls, expected more variety", len(seen))
	}
}
