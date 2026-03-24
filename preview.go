package chroma16

import "fmt"

// slotNames maps terminal color slot indices 0–15 to their conventional names.
var slotNames = [16]string{
	"Black", "Red", "Green", "Yellow",
	"Blue", "Magenta", "Cyan", "White",
	"Bright Black", "Bright Red", "Bright Green", "Bright Yellow",
	"Bright Blue", "Bright Magenta", "Bright Cyan", "Bright White",
}

// renderPreview prints a color swatch of all 16 palette slots to stdout.
// Each row shows the slot index, a colored block using ANSI true-color escapes,
// the hex value, and the conventional slot name. Text color adapts for
// readability based on background lightness (L > 0.55 → dark text).
func renderPreview(p Palette) {
	fmt.Println()
	for i, c := range p.colors {
		hex := toHex(c)
		h := rgbToHSL(c)

		// Choose a contrasting text color based on the background lightness.
		textCode := "38;2;255;255;255" // white text (for dark backgrounds)
		if h.L > 0.55 {
			textCode = "38;2;0;0;0" // black text (for light backgrounds)
		}

		block := fmt.Sprintf("\033[48;2;%d;%d;%dm\033[%sm  ██████  \033[0m",
			c.R, c.G, c.B, textCode)

		fmt.Printf(" %2d  %s  %s  %s\n", i, block, hex, slotNames[i])
	}
	fmt.Println()
}
