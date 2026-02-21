# Dopo Goto

[Dopo Goto](https://dopogoto.com) terminal music and video player.

15 albums. 30+ hours of music.
Live chat. All in your terminal.

## Install

### AI

Paste this into Claude Code, Cursor, or any AI agent with terminal access:

> Install dopogoto terminal music player from github.com/dangerous-person/dopogoto

### Script (macOS & Linux)

```sh
curl -fsSL https://raw.githubusercontent.com/dangerous-person/dopogoto/main/install.sh | sh
```

This downloads the latest binary, installs it to `/usr/local/bin`, and removes the macOS quarantine attribute automatically.

### Download from GitHub Releases

Grab the latest archive from the [Releases page](https://github.com/dangerous-person/dopogoto/releases) and extract it:

```sh
tar -xzf dopogoto_*.tar.gz
./dopogoto
```

**macOS Gatekeeper note:** Since the binary isn't code-signed, macOS will block it. Remove the quarantine attribute after extracting:

```sh
xattr -d com.apple.quarantine dopogoto
```

Or: System Settings → Privacy & Security → scroll down → click "Allow Anyway".

### Build from source

Requires Go 1.25+:

```sh
go install github.com/dangerous-person/dopogoto@latest
```

Or clone and build:

```sh
git clone https://github.com/dangerous-person/dopogoto.git
cd dopogoto
go build -o dopogoto .
./dopogoto
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
- Media/assets are licensed separately (see `ASSETS_LICENSE.md`) and are not covered by MIT.

## Links
https://www.youtube.com/@dopogoto
https://dopogoto.bandcamp.com
https://dopogoto.com
