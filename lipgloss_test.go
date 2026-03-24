//go:build lipgloss

package chroma16

import (
	"testing"
)

func TestLipglossTheme(t *testing.T) {
	p, _ := From("#FF0000") // Red

	theme := p.ToLipglossTheme()

	// Just a quick check to make sure the types map correctly
	// Black is slot 0, so it should be the darkest variation
	h := p.Hex()
	if string(theme.Black) != h[0] {
		t.Errorf("LipglossTheme mapping failed for Black: got %q, want %q", theme.Black, h[0])
	}
	if string(theme.BrightWhite) != h[15] {
		t.Errorf("LipglossTheme mapping failed for BrightWhite: got %q, want %q", theme.BrightWhite, h[15])
	}
}
