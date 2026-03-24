package chroma16

import "strings"

// From generates a 16-color terminal palette from a seed value.
//
// The seed can be:
//   - A hex color string: "#FF6B35" or "FF6B35"
//   - A CSS named color:  "crimson", "steelblue", "gold", etc. (148 names)
//   - Any other string:   deterministically hashed to a hex color via FNV-1a
//
// Returns an error only when the seed starts with "#" or is a bare 6-char
// all-hex string but contains invalid hex characters.
func From(seed string) (Palette, error) {
	hex, err := resolveHex(seed)
	if err != nil {
		return Palette{}, err
	}
	return buildPalette(hex, Neutral, Medium)
}

// New returns a Builder for constructing a palette with custom options.
// Call methods on the builder to configure the palette, then call Build()
// to produce the final Palette.
//
// Example:
//
//	palette, err := chroma16.New().
//	    Seed("#FF6B35").
//	    Mood(chroma16.Warm).
//	    Contrast(chroma16.High).
//	    Build()
func New() *Builder {
	return &Builder{
		mood:     Neutral,
		contrast: Medium,
	}
}

// resolveHex returns a valid "#RRGGBB" hex string from seed.
//
// Disambiguation rules (in order):
//  1. If seed starts with '#' → validate strictly as hex; error on bad chars.
//  2. If seed is exactly 6 all-hex characters → treat as bare RRGGBB hex.
//  3. If seed matches a CSS named color (case-insensitive) → use its hex.
//  4. Everything else → deterministically hash with FNV-1a.
func resolveHex(seed string) (string, error) {
	if strings.HasPrefix(seed, "#") {
		// Explicit hex intent — validate strictly.
		_, err := parseHex(seed)
		if err != nil {
			return "", err
		}
		return seed, nil
	}
	// Bare 6-char string: only treat as hex if ALL characters are hex digits.
	if len(seed) == 6 && isAllHex(seed) {
		_, err := parseHex(seed)
		if err != nil {
			return "", err
		}
		return "#" + strings.ToUpper(seed), nil
	}
	// CSS named color lookup (case-insensitive).
	if h, ok := namedColor(seed); ok {
		return h, nil
	}
	// Anything else: hash to a consistent hex color.
	return hashStringToHex(seed), nil
}

// isAllHex reports whether every byte of s is a valid hexadecimal digit.
func isAllHex(s string) bool {
	for _, c := range s {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}

// buildPalette is the internal factory that converts a validated "#RRGGBB"
// seed into a full Palette using the given Mood and Contrast settings.
func buildPalette(hexSeed string, m Mood, c Contrast) (Palette, error) {
	base, err := parseHex(hexSeed)
	if err != nil {
		return Palette{}, err
	}
	seedHSL := rgbToHSL(base)
	hslColors := generateColors(seedHSL, m, c)
	var p Palette
	for i, h := range hslColors {
		p.colors[i] = hslToRGB(h)
	}
	return p, nil
}
