package chroma16

import (
	"fmt"
	"hash/fnv"
	"math"
	"strconv"
	"strings"
)

// rgb holds a single color as red, green, blue components (0–255 each).
type rgb struct {
	R, G, B uint8
}

// hsl holds a single color as hue (0–360), saturation (0–1), lightness (0–1).
type hsl struct {
	H float64 // 0–360
	S float64 // 0–1
	L float64 // 0–1
}

// parseHex accepts "#RRGGBB" or "RRGGBB" and returns the corresponding rgb.
func parseHex(s string) (rgb, error) {
	s = strings.TrimPrefix(s, "#")
	if len(s) != 6 {
		return rgb{}, fmt.Errorf("chroma16: invalid hex color %q: must be #RRGGBB", "#"+s)
	}
	n, err := strconv.ParseUint(s, 16, 32)
	if err != nil {
		return rgb{}, fmt.Errorf("chroma16: invalid hex color %q: %w", "#"+s, err)
	}
	return rgb{
		R: uint8(n >> 16),
		G: uint8(n >> 8),
		B: uint8(n),
	}, nil
}

// toHex returns the color as a "#RRGGBB" hex string.
func toHex(c rgb) string {
	return fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
}

// rgbToHSL converts an rgb value to its hsl representation.
func rgbToHSL(c rgb) hsl {
	r := float64(c.R) / 255.0
	g := float64(c.G) / 255.0
	b := float64(c.B) / 255.0

	maxC := math.Max(r, math.Max(g, b))
	minC := math.Min(r, math.Min(g, b))
	l := (maxC + minC) / 2.0

	if maxC == minC {
		return hsl{H: 0, S: 0, L: l}
	}

	d := maxC - minC
	s := d / (1.0 - math.Abs(2.0*l-1.0))

	var h float64
	switch maxC {
	case r:
		h = math.Mod((g-b)/d, 6)
	case g:
		h = (b-r)/d + 2
	case b:
		h = (r-g)/d + 4
	}
	h *= 60
	if h < 0 {
		h += 360
	}

	return hsl{H: h, S: s, L: l}
}

// hslToRGB converts an hsl value to its rgb representation.
func hslToRGB(c hsl) rgb {
	h := c.H
	s := c.S
	l := c.L

	chroma := (1.0 - math.Abs(2.0*l-1.0)) * s
	x := chroma * (1.0 - math.Abs(math.Mod(h/60.0, 2.0)-1.0))
	m := l - chroma/2.0

	var r, g, b float64
	switch int(math.Floor(h / 60.0)) {
	case 0:
		r, g, b = chroma, x, 0
	case 1:
		r, g, b = x, chroma, 0
	case 2:
		r, g, b = 0, chroma, x
	case 3:
		r, g, b = 0, x, chroma
	case 4:
		r, g, b = x, 0, chroma
	case 5:
		r, g, b = chroma, 0, x
	}

	return rgb{
		R: uint8(math.Round((r + m) * 255)),
		G: uint8(math.Round((g + m) * 255)),
		B: uint8(math.Round((b + m) * 255)),
	}
}

// clamp returns v clamped to [lo, hi].
func clamp(v, lo, hi float64) float64 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}

// rotateHue rotates hue h by delta degrees, wrapping to [0, 360).
func rotateHue(h, delta float64) float64 {
	return math.Mod(h+delta+360, 360)
}

// hashStringToHex deterministically converts any string to a "#RRGGBB" hex
// color using FNV-1a hashing.
func hashStringToHex(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	n := h.Sum32()
	r := uint8((n >> 16) & 0xFF)
	g := uint8((n >> 8) & 0xFF)
	b := uint8(n & 0xFF)
	return fmt.Sprintf("#%02X%02X%02X", r, g, b)
}

// generateColors produces exactly 16 HSL color values mapped to terminal slots
// 0–15, derived from the given seed color with the requested Mood and Contrast.
func generateColors(seed hsl, m Mood, c Contrast) []hsl {
	adj := moodParams(m)
	normalL, brightL := contrastSpread(c)

	// Guard against degenerate seeds
	seedH := seed.H
	seedS := clamp(seed.S, 0.05, 1.0)
	seedL := clamp(seed.L, 0.05, 0.95)

	// For achromatic seeds (pure gray), use a neutral blue as anchor hue.
	if seed.S < 0.02 {
		seedH = 210
	}

	// Hue offsets for each slot pair (normal[i] and bright[i+8] share the same hue)
	hueOffsets := [8]float64{180, 0, 120, 60, 240, 300, 180, 30}

	// Saturation levels per slot (normals are less saturated)
	normalSat := [8]float64{0.10, 0.75, 0.65, 0.70, 0.65, 0.70, 0.60, 0.15}
	brightSat := [8]float64{0.15, 0.90, 0.80, 0.85, 0.80, 0.85, 0.75, 0.10}

	// Lightness distribution across 8 normal/bright slots
	normalLights := [8]float64{
		normalL - 0.10, // 0 Black
		normalL + 0.02, // 1 Red
		normalL,        // 2 Green
		normalL + 0.04, // 3 Yellow
		normalL - 0.02, // 4 Blue
		normalL + 0.02, // 5 Magenta
		normalL + 0.04, // 6 Cyan
		normalL + 0.38, // 7 White
	}
	brightLights := [8]float64{
		brightL - 0.10, // 8 Bright Black
		brightL + 0.02, // 9 Bright Red
		brightL,        // 10 Bright Green
		brightL + 0.04, // 11 Bright Yellow
		brightL - 0.02, // 12 Bright Blue
		brightL + 0.02, // 13 Bright Magenta
		brightL + 0.04, // 14 Bright Cyan
		brightL + 0.28, // 15 Bright White
	}

	_ = seedL // seedL informs hue choices but we use table-driven L values

	colors := make([]hsl, 16)
	for i := 0; i < 8; i++ {
		h := rotateHue(seedH+adj.HueShift, hueOffsets[i])

		nSat := clamp(normalSat[i]*seedS*adj.SatMul, 0, 1)
		bSat := clamp(brightSat[i]*seedS*adj.SatMul, 0, 1)

		// Slot 0 and 7 are near-achromatic (black / white analogues)
		if i == 0 || i == 7 {
			nSat = clamp(normalSat[i]*adj.SatMul, 0, 1)
			bSat = clamp(brightSat[i]*adj.SatMul, 0, 1)
		}

		nL := clamp(normalLights[i]+adj.LightOff, 0.04, 0.97)
		bL := clamp(brightLights[i]+adj.LightOff, 0.04, 0.97)

		colors[i] = hsl{H: h, S: nSat, L: nL}
		colors[i+8] = hsl{H: h, S: bSat, L: bL}
	}

	return colors
}
