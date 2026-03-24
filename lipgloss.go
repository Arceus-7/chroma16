//go:build lipgloss

package chroma16

import "github.com/charmbracelet/lipgloss"

// LipglossTheme maps all 16 terminal color slots to lipgloss.Color values,
// ready to use with charmbracelet/lipgloss styles.
//
// This type is only available when building with -tags lipgloss:
//
//	go build -tags lipgloss ./...
type LipglossTheme struct {
	Black         lipgloss.Color
	Red           lipgloss.Color
	Green         lipgloss.Color
	Yellow        lipgloss.Color
	Blue          lipgloss.Color
	Magenta       lipgloss.Color
	Cyan          lipgloss.Color
	White         lipgloss.Color
	BrightBlack   lipgloss.Color
	BrightRed     lipgloss.Color
	BrightGreen   lipgloss.Color
	BrightYellow  lipgloss.Color
	BrightBlue    lipgloss.Color
	BrightMagenta lipgloss.Color
	BrightCyan    lipgloss.Color
	BrightWhite   lipgloss.Color
}

// ToLipglossTheme converts the palette into a LipglossTheme whose fields
// correspond to the 16 standard terminal color slots.
//
// This method is only available when building with -tags lipgloss.
// Add lipgloss to your module first:
//
//	go get github.com/charmbracelet/lipgloss
//
// Then build or run with:
//
//	go run -tags lipgloss ./yourprogram
func (p Palette) ToLipglossTheme() LipglossTheme {
	h := p.Hex()
	return LipglossTheme{
		Black:         lipgloss.Color(h[0]),
		Red:           lipgloss.Color(h[1]),
		Green:         lipgloss.Color(h[2]),
		Yellow:        lipgloss.Color(h[3]),
		Blue:          lipgloss.Color(h[4]),
		Magenta:       lipgloss.Color(h[5]),
		Cyan:          lipgloss.Color(h[6]),
		White:         lipgloss.Color(h[7]),
		BrightBlack:   lipgloss.Color(h[8]),
		BrightRed:     lipgloss.Color(h[9]),
		BrightGreen:   lipgloss.Color(h[10]),
		BrightYellow:  lipgloss.Color(h[11]),
		BrightBlue:    lipgloss.Color(h[12]),
		BrightMagenta: lipgloss.Color(h[13]),
		BrightCyan:    lipgloss.Color(h[14]),
		BrightWhite:   lipgloss.Color(h[15]),
	}
}
