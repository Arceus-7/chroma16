package chroma16

import "fmt"

// Builder constructs a Palette with custom options using a fluent API.
// Create one with chroma16.New(). Zero-value Builder is not valid; always
// use New() to ensure sensible defaults.
type Builder struct {
	seed     string
	mood     Mood
	contrast Contrast
}

// Seed sets the seed color. Accepts a "#RRGGBB" hex string, a bare "RRGGBB"
// hex string, or any arbitrary string (which is deterministically hashed to a
// color via FNV-1a).
func (b *Builder) Seed(seed string) *Builder {
	b.seed = seed
	return b
}

// Mood sets the palette mood (temperature and saturation feel).
// Defaults to chroma16.Neutral if not called.
func (b *Builder) Mood(m Mood) *Builder {
	b.mood = m
	return b
}

// Contrast sets the luminance spread of the palette.
// Defaults to chroma16.Medium if not called.
func (b *Builder) Contrast(c Contrast) *Builder {
	b.contrast = c
	return b
}

// Build generates and returns the final Palette based on the configured options.
// Returns an error if no seed was set or if the seed is an invalid hex string.
func (b *Builder) Build() (Palette, error) {
	if b.seed == "" {
		return Palette{}, fmt.Errorf("chroma16: no seed set — call .Seed() before .Build()")
	}
	hex, err := resolveHex(b.seed)
	if err != nil {
		return Palette{}, err
	}
	return buildPalette(hex, b.mood, b.contrast)
}
