package chroma16

import (
	"strings"
	"testing"
)

// ---------------------------------------------------------------------------
// parseHex
// ---------------------------------------------------------------------------

func TestParseHex(t *testing.T) {
	tests := []struct {
		input   string
		wantR   uint8
		wantG   uint8
		wantB   uint8
		wantErr bool
	}{
		{"#FF6B35", 255, 107, 53, false},
		{"FF6B35", 255, 107, 53, false}, // without leading #
		{"#000000", 0, 0, 0, false},
		{"#FFFFFF", 255, 255, 255, false},
		{"#ff0000", 255, 0, 0, false},  // lowercase
		{"#GG0000", 0, 0, 0, true},     // invalid hex char
		{"#FF00", 0, 0, 0, true},       // too short
		{"", 0, 0, 0, true},            // empty
		{"notacolor12", 0, 0, 0, true}, // wrong length even without #
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := parseHex(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("parseHex(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			}
			if !tt.wantErr {
				if got.R != tt.wantR || got.G != tt.wantG || got.B != tt.wantB {
					t.Errorf("parseHex(%q) = {%d,%d,%d}, want {%d,%d,%d}",
						tt.input, got.R, got.G, got.B, tt.wantR, tt.wantG, tt.wantB)
				}
			}
		})
	}
}

// ---------------------------------------------------------------------------
// RGB ↔ HSL round-trip
// ---------------------------------------------------------------------------

func TestRGBRoundTrip(t *testing.T) {
	samples := []rgb{
		{255, 0, 0},   // red
		{0, 255, 0},   // green
		{0, 0, 255},   // blue
		{255, 255, 0}, // yellow
		{0, 255, 255}, // cyan
		{255, 0, 255}, // magenta
		{255, 107, 53},
		{46, 134, 171},
		{128, 64, 192},
		{200, 200, 50},
		{30, 30, 30},
		{220, 220, 220},
	}

	for _, orig := range samples {
		recovered := hslToRGB(rgbToHSL(orig))
		diff := func(a, b uint8) int {
			if a > b {
				return int(a - b)
			}
			return int(b - a)
		}
		if diff(orig.R, recovered.R) > 1 || diff(orig.G, recovered.G) > 1 || diff(orig.B, recovered.B) > 1 {
			t.Errorf("round-trip {%d,%d,%d} → HSL → {%d,%d,%d}: drift > 1",
				orig.R, orig.G, orig.B, recovered.R, recovered.G, recovered.B)
		}
	}
}

// ---------------------------------------------------------------------------
// Known HSL values
// ---------------------------------------------------------------------------

func TestKnownHSL(t *testing.T) {
	tests := []struct {
		name  string
		input rgb
		wantH float64
		wantS float64
		wantL float64
	}{
		{"black", rgb{0, 0, 0}, 0, 0, 0},
		{"white", rgb{255, 255, 255}, 0, 0, 1},
		{"red", rgb{255, 0, 0}, 0, 1, 0.5},
		{"pure green", rgb{0, 255, 0}, 120, 1, 0.5},
		{"pure blue", rgb{0, 0, 255}, 240, 1, 0.5},
	}

	eps := 0.005
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rgbToHSL(tt.input)
			if abs(got.H-tt.wantH) > eps || abs(got.S-tt.wantS) > eps || abs(got.L-tt.wantL) > eps {
				t.Errorf("rgbToHSL(%v) = {%.3f, %.3f, %.3f}, want {%.3f, %.3f, %.3f}",
					tt.input, got.H, got.S, got.L, tt.wantH, tt.wantS, tt.wantL)
			}
		})
	}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// ---------------------------------------------------------------------------
// rotateHue
// ---------------------------------------------------------------------------

func TestRotateHue(t *testing.T) {
	tests := []struct {
		h, delta, want float64
	}{
		{350, 20, 10},
		{0, 0, 0},
		{180, 180, 0},
		{90, 270, 0},
		{30, -60, 330},
	}
	for _, tt := range tests {
		got := rotateHue(tt.h, tt.delta)
		if abs(got-tt.want) > 0.001 {
			t.Errorf("rotateHue(%.1f, %.1f) = %.3f, want %.3f", tt.h, tt.delta, got, tt.want)
		}
	}
}

// ---------------------------------------------------------------------------
// hashStringToHex
// ---------------------------------------------------------------------------

func TestHashStringToHex(t *testing.T) {
	// Deterministic: same input always same output
	for i := 0; i < 5; i++ {
		got := hashStringToHex("ocean")
		if got != hashStringToHex("ocean") {
			t.Error("hashStringToHex is not deterministic")
		}
	}

	// Different inputs produce different outputs (in practice)
	if hashStringToHex("ocean") == hashStringToHex("forest") {
		t.Error("hashStringToHex collision: ocean == forest")
	}

	// Result is a valid #RRGGBB
	h := hashStringToHex("test")
	if len(h) != 7 || h[0] != '#' {
		t.Errorf("hashStringToHex returned invalid format: %q", h)
	}

	// Empty string produces a stable result (FNV of "" is defined)
	e1 := hashStringToHex("")
	e2 := hashStringToHex("")
	if e1 != e2 {
		t.Error("hashStringToHex(\"\") is not stable")
	}
}

// ---------------------------------------------------------------------------
// From()
// ---------------------------------------------------------------------------

func TestFrom(t *testing.T) {
	// Valid hex — no error
	p, err := From("#FF6B35")
	if err != nil {
		t.Fatalf("From(#FF6B35) unexpected error: %v", err)
	}
	if len(p.Hex()) != 16 {
		t.Errorf("expected 16 colors, got %d", len(p.Hex()))
	}

	// Without # — no error
	_, err = From("FF6B35")
	if err != nil {
		t.Fatalf("From(FF6B35) unexpected error: %v", err)
	}

	// Invalid hex-looking input — error
	_, err = From("#GGGGGG")
	if err == nil {
		t.Error("From(#GGGGGG) expected error, got nil")
	}

	// String seed — no error (hashed)
	_, err = From("ocean")
	if err != nil {
		t.Fatalf("From(\"ocean\") unexpected error: %v", err)
	}
}

// ---------------------------------------------------------------------------
// Determinism
// ---------------------------------------------------------------------------

func TestFromDeterministic(t *testing.T) {
	p1, _ := From("ocean")
	p2, _ := From("ocean")
	h1, h2 := p1.Hex(), p2.Hex()
	for i := range h1 {
		if h1[i] != h2[i] {
			t.Errorf("slot %d: %q != %q — not deterministic", i, h1[i], h2[i])
		}
	}
}

// ---------------------------------------------------------------------------
// Palette length
// ---------------------------------------------------------------------------

func TestPaletteLen(t *testing.T) {
	seeds := []string{"#FF6B35", "#000000", "#FFFFFF", "ocean", "forest", "sunset"}
	for _, s := range seeds {
		p, err := From(s)
		if err != nil {
			t.Fatalf("From(%q) error: %v", s, err)
		}
		if n := len(p.Hex()); n != 16 {
			t.Errorf("From(%q).Hex() len = %d, want 16", s, n)
		}
		if n := len(p.ANSI()); n != 16 {
			t.Errorf("From(%q).ANSI() len = %d, want 16", s, n)
		}
		if n := len(p.RGB()); n != 16 {
			t.Errorf("From(%q).RGB() len = %d, want 16", s, n)
		}
	}
}

// ---------------------------------------------------------------------------
// Builder
// ---------------------------------------------------------------------------

func TestBuilderChain(t *testing.T) {
	p, err := New().Seed("#FF6B35").Mood(Warm).Contrast(High).Build()
	if err != nil {
		t.Fatalf("builder chain error: %v", err)
	}
	if len(p.Hex()) != 16 {
		t.Errorf("expected 16 colors, got %d", len(p.Hex()))
	}
}

func TestBuilderNoSeed(t *testing.T) {
	_, err := New().Build()
	if err == nil {
		t.Error("Build() without seed expected error, got nil")
	}
}

// ---------------------------------------------------------------------------
// At()
// ---------------------------------------------------------------------------

func TestPaletteAt(t *testing.T) {
	p, _ := From("#FF6B35")

	for i := 0; i <= 15; i++ {
		s, err := p.At(i)
		if err != nil {
			t.Errorf("At(%d) unexpected error: %v", i, err)
		}
		if len(s) != 7 || s[0] != '#' {
			t.Errorf("At(%d) = %q, want #RRGGBB", i, s)
		}
	}

	_, err := p.At(16)
	if err == nil {
		t.Error("At(16) expected error, got nil")
	}
	_, err = p.At(-1)
	if err == nil {
		t.Error("At(-1) expected error, got nil")
	}
}

// ---------------------------------------------------------------------------
// All Moods & Contrasts
// ---------------------------------------------------------------------------

func TestAllMoods(t *testing.T) {
	moods := []Mood{Neutral, Warm, Cool, Dark, Pastel, Neon}
	for _, m := range moods {
		p, err := New().Seed("#2E86AB").Mood(m).Build()
		if err != nil {
			t.Errorf("Mood %d: unexpected error: %v", m, err)
		}
		if len(p.Hex()) != 16 {
			t.Errorf("Mood %d: expected 16 colors", m)
		}
	}
}

func TestAllContrasts(t *testing.T) {
	contrasts := []Contrast{Medium, High, Low}
	for _, c := range contrasts {
		p, err := New().Seed("#2E86AB").Contrast(c).Build()
		if err != nil {
			t.Errorf("Contrast %d: unexpected error: %v", c, err)
		}
		if len(p.Hex()) != 16 {
			t.Errorf("Contrast %d: expected 16 colors", c)
		}
	}
}

// ---------------------------------------------------------------------------
// Edge cases: Black and White seeds
// ---------------------------------------------------------------------------

func TestEdgeCaseSeedBlack(t *testing.T) {
	p, err := From("#000000")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	hexes := p.Hex()
	// Not all identical
	allSame := true
	for _, h := range hexes[1:] {
		if h != hexes[0] {
			allSame = false
			break
		}
	}
	if allSame {
		t.Error("black seed produced all-identical colors — no visible contrast")
	}
}

func TestEdgeCaseSeedWhite(t *testing.T) {
	p, err := From("#FFFFFF")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	hexes := p.Hex()
	allSame := true
	for _, h := range hexes[1:] {
		if h != hexes[0] {
			allSame = false
			break
		}
	}
	if allSame {
		t.Error("white seed produced all-identical colors — no visible contrast")
	}
}

// ---------------------------------------------------------------------------
// ANSI format sanity
// ---------------------------------------------------------------------------

func TestANSIFormat(t *testing.T) {
	p, _ := From("#FF6B35")
	for i, seq := range p.ANSI() {
		if !strings.HasPrefix(seq, "\033]4;") {
			t.Errorf("ANSI[%d] does not start with \\033]4;: %q", i, seq)
		}
	}
}

// ---------------------------------------------------------------------------
// Hex format sanity
// ---------------------------------------------------------------------------

func TestHexFormat(t *testing.T) {
	p, _ := From("sunset")
	for i, h := range p.Hex() {
		if len(h) != 7 || h[0] != '#' {
			t.Errorf("Hex[%d] = %q, not #RRGGBB format", i, h)
		}
	}
}
