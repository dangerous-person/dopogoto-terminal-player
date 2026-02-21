package panels

// Theme holds all color values (256-color codes) for the UI.
type Theme struct {
	Name string

	// Background
	Bg string

	// Border colors (inactive)
	BorderColor string
	FadeColor   string
	CornerColor string

	// Border colors (active/focused)
	ActiveBorderColor string
	ActiveFadeColor   string
	ActiveCornerColor string

	// Title gradient (3 colors: brightest to dimmest)
	TitleGrad1 string
	TitleGrad2 string
	TitleGrad3 string

	// Selection
	SelectionBg string
	SelectionFg string // text on selection bg ("" defaults to "231")

	// Text
	TextColor      string // body text (album names, chat messages)
	TextDim        string // dimmed body text (track numbers)
	ChatNameColor  string // accent (nicknames, selected items)
	ChatInputColor string // chat input text + cursor ("" defaults to ChatNameColor)
	ChatOffline    string

	// Help bar
	HelpBracket string
	HelpKey     string
	HelpLabel   string // help bar label color ("" defaults to TextColor)

	// Controls
	PlayedDefault string // default played portion color
	PlayedBg      string // timeline played bg ("" defaults to SelectionBg)
	UnplayedColor string

	// Album colors
	AlbumColors []string

	// Video tint (0 = normal; applies monochrome phosphor shader)
	VideoTintHue float64 // hue 0-360
	VideoTintSat float64 // saturation 0-100
}

// ThemeMono is a monochrome theme — white/gray on black.
var ThemeMono = Theme{
	Name:              "Mono",
	Bg:                "16",
	BorderColor:       "238",
	FadeColor:         "243",
	CornerColor:       "248",
	ActiveBorderColor: "250",
	ActiveFadeColor:   "253",
	ActiveCornerColor: "231",
	TitleGrad1:        "231",
	TitleGrad2:        "253",
	TitleGrad3:        "250",
	SelectionBg:       "235",
	TextColor:         "250",
	TextDim:           "240",
	ChatNameColor:     "231",
	ChatOffline:       "240",
	HelpBracket:       "240",
	HelpKey:           "231",
	PlayedDefault:     "250",
	UnplayedColor:     "235",
	AlbumColors:       []string{"231", "255", "254", "253", "252", "251", "250", "249", "248", "247", "246", "245", "244", "243", "242"},
}

// ThemeMonoColor — Mono palette with color video, inverted selection.
var ThemeMonoColor = Theme{
	Name:              "mono-plus",
	Bg:                "16",
	BorderColor:       "238",
	FadeColor:         "243",
	CornerColor:       "248",
	ActiveBorderColor: "250",
	ActiveFadeColor:   "253",
	ActiveCornerColor: "231",
	TitleGrad1:        "231",
	TitleGrad2:        "253",
	TitleGrad3:        "250",
	SelectionBg:       "220",
	SelectionFg:       "16",
	PlayedBg:          "235",
	TextColor:         "250",
	TextDim:           "240",
	ChatNameColor:     "196",
	ChatOffline:       "240",
	HelpBracket:       "240",
	HelpKey:           "220",
	PlayedDefault:     "250",
	UnplayedColor:     "235",
	AlbumColors:       []string{"231", "255", "254", "253", "252", "251", "250", "249", "248", "247", "246", "245", "244", "243", "242"},
}

// ThemeMonoPink — Mono palette with pink/lavender accents.
var ThemeMonoPink = Theme{
	Name:              "mono-pink",
	Bg:                "16",
	BorderColor:       "238",
	FadeColor:         "243",
	CornerColor:       "248",
	ActiveBorderColor: "250",
	ActiveFadeColor:   "253",
	ActiveCornerColor: "231",
	TitleGrad1:        "231",
	TitleGrad2:        "253",
	TitleGrad3:        "250",
	SelectionBg:       "218",
	SelectionFg:       "16",
	PlayedBg:          "235",
	TextColor:         "250",
	TextDim:           "240",
	ChatNameColor:     "141",
	ChatOffline:       "240",
	HelpBracket:       "240",
	HelpKey:           "218",
	PlayedDefault:     "250",
	UnplayedColor:     "235",
	AlbumColors:       []string{"231", "255", "254", "253", "252", "251", "250", "249", "248", "247", "246", "245", "244", "243", "242"},
}

// ThemeMonoViolet — Mono palette with blue-purple accents on dark navy.
var ThemeMonoViolet = Theme{
	Name:              "mono-violet",
	Bg:                "17",
	BorderColor:       "238",
	FadeColor:         "243",
	CornerColor:       "248",
	ActiveBorderColor: "250",
	ActiveFadeColor:   "253",
	ActiveCornerColor: "231",
	TitleGrad1:        "231",
	TitleGrad2:        "253",
	TitleGrad3:        "250",
	SelectionBg:       "135",
	SelectionFg:       "",
	PlayedBg:          "63",
	TextColor:         "250",
	TextDim:           "240",
	ChatNameColor:     "63",
	ChatOffline:       "240",
	HelpBracket:       "240",
	HelpKey:           "135",
	PlayedDefault:     "250",
	UnplayedColor:     "63",
	AlbumColors:       []string{"231", "255", "254", "253", "252", "251", "250", "249", "248", "247", "246", "245", "244", "243", "242"},
}

// ThemeMonoBlush — Mono palette with pink/blue accents on dark blue.
var ThemeMonoBlush = Theme{
	Name:              "mono-blush",
	Bg:                "18",
	BorderColor:       "238",
	FadeColor:         "243",
	CornerColor:       "248",
	ActiveBorderColor: "250",
	ActiveFadeColor:   "253",
	ActiveCornerColor: "231",
	TitleGrad1:        "231",
	TitleGrad2:        "253",
	TitleGrad3:        "250",
	SelectionBg:       "175",
	SelectionFg:       "",
	PlayedBg:          "17",
	TextColor:         "250",
	TextDim:           "246",
	ChatNameColor:     "51",
	ChatOffline:       "240",
	HelpBracket:       "240",
	HelpKey:           "175",
	HelpLabel:         "231",
	PlayedDefault:     "250",
	UnplayedColor:     "17",
	AlbumColors:       []string{"231", "255", "254", "253", "252", "251", "250", "249", "248", "247", "246", "245", "244", "243", "242"},
}

// ThemeMonoMint — Mono palette with yellow-green/green accents on dark green.
var ThemeMonoMint = Theme{
	Name:              "mono-mint",
	Bg:                "233",
	BorderColor:       "238",
	FadeColor:         "243",
	CornerColor:       "248",
	ActiveBorderColor: "250",
	ActiveFadeColor:   "253",
	ActiveCornerColor: "231",
	TitleGrad1:        "231",
	TitleGrad2:        "253",
	TitleGrad3:        "250",
	SelectionBg:       "192",
	SelectionFg:       "16",
	PlayedBg:          "235",
	TextColor:         "250",
	TextDim:           "246",
	ChatNameColor:     "78",
	ChatOffline:       "240",
	HelpBracket:       "78",
	HelpKey:           "192",
	HelpLabel:         "231",
	PlayedDefault:     "250",
	UnplayedColor:     "235",
	AlbumColors:       []string{"231", "255", "254", "253", "252", "251", "250", "249", "248", "247", "246", "245", "244", "243", "242"},
	VideoTintHue:      100,
	VideoTintSat:      70,
}

// ThemeMonoEmber — Mono palette with orange/yellow accents on dark warm bg.
var ThemeMonoEmber = Theme{
	Name:              "mono-ember",
	Bg:                "232",
	BorderColor:       "238",
	FadeColor:         "243",
	CornerColor:       "248",
	ActiveBorderColor: "250",
	ActiveFadeColor:   "253",
	ActiveCornerColor: "231",
	TitleGrad1:        "231",
	TitleGrad2:        "253",
	TitleGrad3:        "250",
	SelectionBg:       "220",
	SelectionFg:       "16",
	TextColor:         "250",
	TextDim:           "246",
	ChatNameColor:     "208",
	ChatInputColor:    "220",
	ChatOffline:       "240",
	HelpBracket:       "208",
	HelpKey:           "220",
	HelpLabel:         "231",
	PlayedDefault:     "250",
	PlayedBg:          "233",
	UnplayedColor:     "233",
	AlbumColors:       []string{"231", "255", "254", "253", "252", "251", "250", "249", "248", "247", "246", "245", "244", "243", "242"},
	VideoTintHue:      20,
	VideoTintSat:      45,
}

var themes = []Theme{ThemeMonoColor, ThemeMonoPink, ThemeMonoViolet, ThemeMonoBlush, ThemeMonoMint, ThemeMonoEmber, ThemeMono}
var themeIdx = 0

// CurrentTheme returns the active theme.
func CurrentTheme() *Theme {
	return &themes[themeIdx]
}

// CycleTheme advances to the next theme and returns it.
func CycleTheme() *Theme {
	themeIdx = (themeIdx + 1) % len(themes)
	return &themes[themeIdx]
}
