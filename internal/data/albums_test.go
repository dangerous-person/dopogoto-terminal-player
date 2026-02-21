package data

import (
	"testing"
	"time"
)

func TestFormatDuration(t *testing.T) {
	tests := []struct {
		name string
		d    time.Duration
		want string
	}{
		{"zero", 0, "0:00"},
		{"seconds only", 45 * time.Second, "0:45"},
		{"minutes and seconds", 3*time.Minute + 27*time.Second, "3:27"},
		{"exact minute", 5 * time.Minute, "5:00"},
		{"over an hour", 1*time.Hour + 2*time.Minute + 3*time.Second, "1:02:03"},
		{"typical track", 4*time.Minute + 15*time.Second, "4:15"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormatDuration(tt.d)
			if got != tt.want {
				t.Errorf("FormatDuration(%v) = %q, want %q", tt.d, got, tt.want)
			}
		})
	}
}

func TestAlbumCatalog(t *testing.T) {
	if len(Albums) != 15 {
		t.Errorf("expected 15 albums, got %d", len(Albums))
	}

	for i, album := range Albums {
		if album.Title == "" {
			t.Errorf("album %d has empty title", i)
		}
		if len(album.Tracks) == 0 {
			t.Errorf("album %q has no tracks", album.Title)
		}
		for j, track := range album.Tracks {
			if track.Title == "" {
				t.Errorf("album %q track %d has empty title", album.Title, j)
			}
			if track.URL == "" {
				t.Errorf("album %q track %q has empty URL", album.Title, track.Title)
			}
		}
	}
}
