# Dopo Goto

[Dopo Goto](https://dopogoto.com) terminal music and video player.

15 albums. 30+ hours of music. 
Live chat. All in your terminal.


## Install

```sh
go install github.com/dangerous-person/dopogoto@latest
```

Or build from source:

```sh
git clone https://github.com/dangerous-person/dopogoto.git
cd dopogoto
go build -o dopogoto .
./dopogoto
```

## Requirements

- Go 1.25+
- A 256-color terminal (Ghostty, Terminal.app, iTerm2, Windows Terminal, etc.)
- Terminal size 120x40 or larger
- Audio output device

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
