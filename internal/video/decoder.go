package video

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"

	"github.com/andybalholm/brotli"
)

// Cell represents a single character cell in the video buffer.
type Cell struct {
	CharIdx  int
	ColorIdx int
}

// rawVideoData handles both v1 (frames as [][]int) and v2 (frames as []string).
// Chars can be either a string or an array of strings in the JSON.
type rawVideoData struct {
	V       int             `json:"v"`
	W       int             `json:"w"`
	H       int             `json:"h"`
	FPS     int             `json:"fps"`
	Chars   json.RawMessage `json:"chars"`
	Palette []string        `json:"palette"`
	Frames  json.RawMessage `json:"frames"`
}

// VideoData is the normalized format with decoded int frames.
type VideoData struct {
	W       int
	H       int
	FPS     int
	Chars   string
	Palette []string
	Frames  [][]int
}

// Decoder handles keyframe/delta decoding of ascii-term video format.
type Decoder struct {
	Data          VideoData
	Buffer        []Cell
	KeyframeIndex []int
	chars         []rune
}

// NewDecoder creates a decoder from JSON or gzip-compressed JSON data.
func NewDecoder(data []byte) (*Decoder, error) {
	// Try gzip decompression first
	if len(data) >= 2 && data[0] == 0x1f && data[1] == 0x8b {
		gr, err := gzip.NewReader(bytes.NewReader(data))
		if err != nil {
			return nil, fmt.Errorf("gzip open: %w", err)
		}
		data, err = io.ReadAll(gr)
		gr.Close()
		if err != nil {
			return nil, fmt.Errorf("gzip read: %w", err)
		}
	} else {
		// Try brotli decompression (no magic number — try and fall back)
		br := brotli.NewReader(bytes.NewReader(data))
		decoded, err := io.ReadAll(br)
		if err == nil && len(decoded) > 0 {
			data = decoded
		}
	}

	var raw rawVideoData
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("decode video json: %w", err)
	}

	// Parse chars: can be a string or array of strings
	var chars string
	if len(raw.Chars) > 0 && raw.Chars[0] == '"' {
		// JSON string: "abc"
		if err := json.Unmarshal(raw.Chars, &chars); err != nil {
			return nil, fmt.Errorf("decode chars string: %w", err)
		}
	} else {
		// JSON array: [" ",".",","]
		var charArr []string
		if err := json.Unmarshal(raw.Chars, &charArr); err != nil {
			return nil, fmt.Errorf("decode chars array: %w", err)
		}
		for _, c := range charArr {
			chars += c
		}
	}

	vd := VideoData{
		W:       raw.W,
		H:       raw.H,
		FPS:     raw.FPS,
		Chars:   chars,
		Palette: raw.Palette,
	}

	if raw.V == 2 {
		// v2: frames are base-93 encoded strings
		var strFrames []string
		if err := json.Unmarshal(raw.Frames, &strFrames); err != nil {
			return nil, fmt.Errorf("decode v2 frames: %w", err)
		}
		charCount := len([]rune(chars))
		if charCount == 0 {
			return nil, fmt.Errorf("v2 video has empty chars set")
		}
		vd.Frames = decodeV2Frames(strFrames, charCount)
	} else {
		// v1: frames are arrays of ints
		if err := json.Unmarshal(raw.Frames, &vd.Frames); err != nil {
			return nil, fmt.Errorf("decode v1 frames: %w", err)
		}
	}

	if vd.W == 0 || vd.H == 0 || len(vd.Frames) == 0 {
		return nil, fmt.Errorf("invalid video data: w=%d h=%d frames=%d", vd.W, vd.H, len(vd.Frames))
	}

	buf := make([]Cell, vd.W*vd.H)

	var kfIdx []int
	for i, frame := range vd.Frames {
		if len(frame) > 0 && frame[0] == 1 {
			kfIdx = append(kfIdx, i)
		}
	}

	return &Decoder{
		Data:          vd,
		Buffer:        buf,
		KeyframeIndex: kfIdx,
		chars:         []rune(vd.Chars),
	}, nil
}

// decodeV2Frames converts base-93 string frames to v1 int arrays.
// Ported from play.js lines 34-55.
func decodeV2Frames(strFrames []string, charCount int) [][]int {
	// Build base-93 decode table: printable ASCII 32-126, excluding " (34) and \ (92)
	var decodeTable [128]int
	idx := 0
	for c := 32; c <= 126; c++ {
		if c != 34 && c != 92 {
			decodeTable[c] = idx
			idx++
		}
	}

	frames := make([][]int, len(strFrames))
	for i, s := range strFrames {
		raw := []byte(s) // base-93 uses ASCII only
		if len(raw) == 0 {
			frames[i] = []int{0}
			continue
		}
		isKey := raw[0] == 'K'

		decoded := []int{0}
		if isKey {
			decoded[0] = 1
		}

		for j := 1; j+3 < len(raw); j += 4 {
			// Bounds-check indices into decodeTable
			if raw[j] >= 128 || raw[j+1] >= 128 || raw[j+2] >= 128 || raw[j+3] >= 128 {
				continue
			}
			a := decodeTable[raw[j]]*93 + decodeTable[raw[j+1]]
			b := decodeTable[raw[j+2]]*93 + decodeTable[raw[j+3]]
			if isKey {
				decoded = append(decoded, a%charCount, a/charCount, b)
			} else {
				decoded = append(decoded, a, b%charCount, b/charCount)
			}
		}

		frames[i] = decoded
	}
	return frames
}

// ApplyFrame applies a single frame (keyframe or delta) to the buffer.
func (d *Decoder) ApplyFrame(idx int) {
	if idx < 0 || idx >= len(d.Data.Frames) {
		return
	}

	frame := d.Data.Frames[idx]
	if len(frame) == 0 {
		return
	}

	if frame[0] == 1 {
		// Keyframe: RLE triplets (charIdx, colorIdx, count)
		ci := 0
		for i := 1; i+2 < len(frame); i += 3 {
			charIdx := frame[i]
			colorIdx := frame[i+1]
			count := frame[i+2]
			for j := 0; j < count && ci < len(d.Buffer); j++ {
				d.Buffer[ci].CharIdx = charIdx
				d.Buffer[ci].ColorIdx = colorIdx
				ci++
			}
		}
	} else {
		// Delta: sparse triplets (pos, charIdx, colorIdx)
		for i := 1; i+2 < len(frame); i += 3 {
			pos := frame[i]
			if pos < len(d.Buffer) {
				d.Buffer[pos].CharIdx = frame[i+1]
				d.Buffer[pos].ColorIdx = frame[i+2]
			}
		}
	}
}

// SeekTo seeks to a target frame by finding the nearest keyframe
// and applying all frames forward.
func (d *Decoder) SeekTo(targetIdx int) {
	kf := 0
	for _, ki := range d.KeyframeIndex {
		if ki <= targetIdx {
			kf = ki
		} else {
			break
		}
	}
	for i := kf; i <= targetIdx; i++ {
		d.ApplyFrame(i)
	}
}

// Char returns the character for a given char index.
func (d *Decoder) Char(idx int) rune {
	if idx >= 0 && idx < len(d.chars) {
		return d.chars[idx]
	}
	return ' '
}

func (d *Decoder) Width() int       { return d.Data.W }
func (d *Decoder) Height() int      { return d.Data.H }
func (d *Decoder) FPS() int         { return d.Data.FPS }
func (d *Decoder) TotalFrames() int { return len(d.Data.Frames) }
