package panels

import (
	"fmt"
	"strings"
	"time"

	"github.com/dangerous-person/dopogoto/internal/data"
)

type PlayState int

const (
	StateStopped PlayState = iota
	StateBuffering
	StatePlaying
	StatePaused
)

type Controls struct {
	State      PlayState
	TrackTitle string
	Position   time.Duration
	Duration   time.Duration
	Volume     int // 0-10
	Width      int
	Height     int
	Shuffle    bool
	Repeat     bool
	AlbumColor string // 256-color for played portion of timeline
}

func NewControls() Controls {
	return Controls{
		State:  StateStopped,
		Volume: 7,
	}
}

const controlsH = 3 // fixed height: 2 border + 1 content

// ControlsHeight returns the fixed height of the controls panel.
func ControlsHeight() int {
	return controlsH
}

func (c Controls) View() string {
	t := CurrentTheme()

	w := c.Width
	if w < 20 {
		w = 20
	}
	contentW := w - 4 // border (2) + padding (1 each side)
	if contentW < 16 {
		contentW = 16
	}

	var b strings.Builder

	// Top border with title
	titleAnsi := BuildTitleGradient("Now Playing", t.TitleGrad1, t.TitleGrad2, t.TitleGrad3)
	title := fmt.Sprintf(" %s\x1b[38;5;%sm ", titleAnsi, t.BorderColor)
	titleVisLen := 13
	remaining := contentW + 2 - titleVisLen
	if remaining < 0 {
		remaining = 0
	}
	leftPad := remaining / 2
	rightPad := remaining - leftPad

	b.WriteString(fmt.Sprintf("\x1b[38;5;%sm╭", t.CornerColor))
	b.WriteString(FadeBorder(leftPad, FadeDashes, t.FadeColor, t.BorderColor))
	b.WriteString(title)
	b.WriteString(FadeBorder(rightPad, FadeDashes, t.FadeColor, t.BorderColor))
	b.WriteString(fmt.Sprintf("\x1b[38;5;%sm╮\x1b[0m\n", t.CornerColor))

	// Build display text: "Song Name [00:00/15:38]"
	trackTitle := c.TrackTitle
	if trackTitle == "" {
		trackTitle = "Choose a Song"
	}

	posStr := data.FormatDuration(c.Position)
	durStr := data.FormatDuration(c.Duration)
	timer := fmt.Sprintf("[%s/%s]", posStr, durStr)
	timerLen := len([]rune(timer))

	// Title + space + timer
	maxTitle := contentW - timerLen - 1
	if maxTitle < 5 {
		maxTitle = 5
	}
	trackTitle = truncate(trackTitle, maxTitle)

	display := trackTitle + " " + timer
	displayRunes := []rune(display)
	dLen := len(displayRunes)
	timerStart := len([]rune(trackTitle)) + 1 // index within displayRunes where timer begins
	displayStart := (contentW - dLen) / 2
	if displayStart < 0 {
		displayStart = 0
	}

	var filled int
	if c.Duration > 0 {
		filled = int(float64(contentW) * float64(c.Position) / float64(c.Duration))
	}
	if filled > contentW {
		filled = contentW
	}

	playedFg := c.AlbumColor
	if playedFg == "" {
		playedFg = t.PlayedDefault
	}

	playedBg := t.PlayedBg
	if playedBg == "" {
		playedBg = t.SelectionBg
	}

	var line strings.Builder
	for i := 0; i < contentW; i++ {
		played := i < filled
		inDisplay := i >= displayStart && i < displayStart+dLen

		if inDisplay {
			ch := displayRunes[i-displayStart]
			isTimer := (i - displayStart) >= timerStart
			var fg string
			if isTimer {
				fg = t.TextDim
			} else if played {
				fg = playedFg
			} else {
				fg = t.TextColor
			}
			if played {
				line.WriteString(fmt.Sprintf("\x1b[48;5;%sm\x1b[38;5;%sm%c", playedBg, fg, ch))
			} else {
				line.WriteString(fmt.Sprintf("\x1b[49m\x1b[38;5;%sm%c", fg, ch))
			}
		} else {
			if played {
				line.WriteString(fmt.Sprintf("\x1b[48;5;%sm ", playedBg))
			} else {
				line.WriteString("\x1b[49m ")
			}
		}
	}
	line.WriteString("\x1b[0m")
	writeBorderedLine(&b, t.BorderColor, t.FadeColor, line.String(), contentW, 0, 1, true)

	// Bottom border
	b.WriteString(fmt.Sprintf("\x1b[38;5;%sm╰", t.CornerColor))
	b.WriteString(FadeBorder(contentW+2, FadeDashes, t.FadeColor, t.BorderColor))
	b.WriteString(fmt.Sprintf("\x1b[38;5;%sm╯\x1b[0m", t.CornerColor))

	return b.String()
}
