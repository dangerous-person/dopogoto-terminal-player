# Dopo Goto

[Dopo Goto](https://dopogoto.com) terminal music and video player.

15 albums. 30+ hours of music.
Live chat. All in your terminal.

## Installation

### Ask AI

Paste this into Claude Code, Codex, Cursor, or any AI agent with terminal access:

> Install dopogoto terminal music player from github.com/dangerous-person/dopogoto

### macOS & Linux

```sh
curl -fsSL https://raw.githubusercontent.com/dangerous-person/dopogoto/main/install.sh | sh
```

Or download manually from [Releases](https://github.com/dangerous-person/dopogoto/releases).

**macOS Gatekeeper:** If macOS blocks the binary, remove the quarantine attribute:

```sh
xattr -d com.apple.quarantine dopogoto
```

### Windows

Download `dopogoto_*_windows_amd64.zip` from [Releases](https://github.com/dangerous-person/dopogoto/releases), extract, and run `dopogoto.exe`.

### Build from source

Requires Go 1.25+:

```sh
go install github.com/dangerous-person/dopogoto@latest
```

## Requirements

- A 256-color terminal (Ghostty, Terminal.app, iTerm2, Windows Terminal, etc.)
- Terminal size 120x40 or larger
- Audio output device

**Linux:** ALSA is required for audio playback.
- Debian/Ubuntu: `sudo apt install libasound2-dev`
- Fedora: `sudo dnf install alsa-lib-devel`
- Arch: `sudo pacman -S alsa-lib`

## Controls

| Key | Action |
|-----|--------|
| TAB | Cycle focus: Albums / Tracks / Chat |
| UP/DOWN | Navigate lists |
| ENTER | Play track or send chat message |
| SPACE | Pause / resume |
| N / P | Next / previous track |
| +/- | Volume up / down |
| S | Shuffle |
| R | Repeat |
| T | Change theme |
| LEFT/RIGHT | Seek -/+ 10s |
| Q | Quit |

## Chat

Type a message and press Enter.

- `/nick name` -- set your nickname (saved locally)
- `/reset` -- go anonymous

## Licensing

- Source code is licensed under MIT (see `LICENSE`).
- Media/assets are licensed separately (see [ASSETS_LICENSE.md](ASSETS_LICENSE.md)) and are not covered by MIT.

## Links
- https://www.youtube.com/@dopogoto
- https://dopogoto.bandcamp.com
- https://dopogoto.com
