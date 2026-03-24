package chroma16

import (
	"encoding/json"
	"strings"
	"testing"
)

// ---------------------------------------------------------------------------
// Named colors
// ---------------------------------------------------------------------------

func TestNamedColors(t *testing.T) {
	// Case-insensitive lookup works
	p1, _ := From("CRImson") // #DC143C
	if h, _ := p1.At(0); len(h) != 7 {
		t.Errorf("named color lookup failed")
	}

	// Unknown string hashes instead
	p2, _ := From("notarealcolor123")
	if len(p2.Hex()) != 16 {
		t.Errorf("falling back to hash failed")
	}
}

// ---------------------------------------------------------------------------
// Export functions (Alacritty, Kitty, WT, Xresources)
// ---------------------------------------------------------------------------

func TestExportFormats(t *testing.T) {
	p, _ := From("#000000")

	// Alacritty Check
	ala := p.ToAlacritty()
	if !strings.Contains(ala, "[colors.primary]") || !strings.Contains(ala, "black   = \"#") {
		t.Errorf("ToAlacritty output malformed: \n%s", ala)
	}

	// Kitty Check
	kit := p.ToKitty()
	if !strings.Contains(kit, "color0 ") || !strings.Contains(kit, "background ") {
		t.Errorf("ToKitty output malformed\n")
	}

	// Windows Terminal Check
	wt := p.ToWindowsTerminal("MyTheme")
	if !strings.Contains(wt, "\"name\": \"MyTheme\"") || !strings.Contains(wt, "\"brightPurple\": \"#") {
		t.Errorf("ToWindowsTerminal output malformed")
	}

	// Xresources Check
	xres := p.ToXresources()
	if !strings.Contains(xres, "*.color0 :  #") || !strings.Contains(xres, "*.foreground: #") {
		t.Errorf("ToXresources output malformed")
	}
}

// ---------------------------------------------------------------------------
// Blend, Complement, Analogous
// ---------------------------------------------------------------------------

func TestBlend(t *testing.T) {
	p1, _ := From("#FF0000") // Red
	p2, _ := From("#0000FF") // Blue

	// Blend 50%
	pMid := Blend(p1, p2, 0.5)

	if len(pMid.Hex()) != 16 {
		t.Errorf("Blend lost colors")
	}
}

func TestAnalogousAndComplement(t *testing.T) {
	p, _ := From("#FF6B35")

	pC := p.Complement()
	if p.Hex()[0] == pC.Hex()[0] {
		t.Errorf("Complement did not change palette")
	}

	pA := p.Analogous(60)
	if p.Hex()[0] == pA.Hex()[0] {
		t.Errorf("Analogous did not change palette")
	}
}

// ---------------------------------------------------------------------------
// JSON Round-trip
// ---------------------------------------------------------------------------

func TestJSON(t *testing.T) {
	orig, _ := From("#2E86AB")

	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}

	restored, err := FromJSON(data)
	if err != nil {
		t.Fatalf("FromJSON failed: %v", err)
	}

	if orig.Hex()[0] != restored.Hex()[0] {
		t.Errorf("JSON round-trip failed: want %v, got %v", orig.Hex()[0], restored.Hex()[0])
	}
}

// Test malformed JSON
func TestJSONMalformed(t *testing.T) {
	// Wrong version
	badVer := `{"version":2,"colors":["#000000","#000000","#000000","#000000","#000000","#000000","#000000","#000000","#000000","#000000","#000000","#000000","#000000","#000000","#000000","#000000"]}`
	if _, err := FromJSON([]byte(badVer)); err == nil {
		t.Errorf("FromJSON expected error on wrong version")
	}

	// Wrong length
	badLen := `{"version":1,"colors":["#000000"]}`
	if _, err := FromJSON([]byte(badLen)); err == nil {
		t.Errorf("FromJSON expected error on wrong length")
	}

	// Bad hex
	badHex := `{"version":1,"colors":["#GGGGGG","#000000","#000000","#000000","#000000","#000000","#000000","#000000","#000000","#000000","#000000","#000000","#000000","#000000","#000000","#000000"]}`
	if _, err := FromJSON([]byte(badHex)); err == nil {
		t.Errorf("FromJSON expected error on bad hex")
	}
}
