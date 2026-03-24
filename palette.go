package chroma16

import "fmt"

// Palette holds a 16-color terminal palette derived from a single seed color.
// Colors are ordered to match standard terminal color slot indices 0–15:
// 0=Black, 1=Red, 2=Green, 3=Yellow, 4=Blue, 5=Magenta, 6=Cyan, 7=White,
// 8=Bright Black, 9=Bright Red, … 15=Bright White.
type Palette struct {
	colors [16]rgb
}

// Hex returns the 16 palette colors as hex strings in #RRGGBB format,
// ordered by terminal color slot index (0 = Black, 15 = Bright White).
func (p Palette) Hex() []string {
	out := make([]string, 16)
	for i, c := range p.colors {
		out[i] = toHex(c)
	}
	return out
}

// ANSI returns 16 ANSI OSC 4 escape sequences that set each terminal color
// slot to the palette's corresponding color. Apply all 16 to repaint a
// terminal session with the full palette. The format used is the XTerm OSC 4
// sequence, widely supported in modern terminals.
func (p Palette) ANSI() []string {
	out := make([]string, 16)
	for i, c := range p.colors {
		out[i] = fmt.Sprintf("\033]4;%d;rgb:%02x/%02x/%02x\033\\",
			i, c.R, c.G, c.B)
	}
	return out
}

// RGB returns the 16 palette colors as raw [R, G, B] uint8 triples.
func (p Palette) RGB() [][3]uint8 {
	out := make([][3]uint8, 16)
	for i, c := range p.colors {
		out[i] = [3]uint8{c.R, c.G, c.B}
	}
	return out
}

// At returns the hex color string for the given terminal color slot index.
// Returns an error if index is outside [0, 15].
func (p Palette) At(index int) (string, error) {
	if index < 0 || index > 15 {
		return "", fmt.Errorf("chroma16: index %d out of range [0, 15]", index)
	}
	return toHex(p.colors[index]), nil
}

// Preview prints a visual color swatch of all 16 palette slots to stdout,
// showing the terminal slot index, a colored block, the hex value, and the
// slot name.
func (p Palette) Preview() {
	renderPreview(p)
}
