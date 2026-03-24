package chroma16

import (
	"fmt"
	"math"
)

// Blend returns a new Palette that is the linear interpolation between p1 and
// p2. t=0 returns a copy of p1, t=1 returns a copy of p2. Values outside [0,1]
// are clamped. Hue interpolation takes the shortest arc around the color wheel.
func Blend(p1, p2 Palette, t float64) Palette {
	t = clamp(t, 0, 1)
	var out Palette
	for i := range out.colors {
		h1 := rgbToHSL(p1.colors[i])
		h2 := rgbToHSL(p2.colors[i])

		// Shortest-path hue interpolation to avoid spinning the wrong way.
		dH := h2.H - h1.H
		if dH > 180 {
			dH -= 360
		} else if dH < -180 {
			dH += 360
		}

		blended := hsl{
			H: math.Mod(h1.H+dH*t+360, 360),
			S: h1.S + (h2.S-h1.S)*t,
			L: h1.L + (h2.L-h1.L)*t,
		}
		out.colors[i] = hslToRGB(blended)
	}
	return out
}

// Complement returns a new Palette with all hues rotated 180°, producing the
// complementary color palette.
func (p Palette) Complement() Palette {
	return p.Analogous(180)
}

// Analogous returns a new Palette with all hues rotated by delta degrees.
// Positive delta shifts clockwise; negative shifts counterclockwise.
// The rotation wraps correctly around the 360°/0° boundary.
//
// Example:
//
//	warm := palette.Analogous(-15) // nudge toward warmer hues
//	split := palette.Analogous(30) // shift to an analogous neighbor
func (p Palette) Analogous(delta float64) Palette {
	var out Palette
	for i, c := range p.colors {
		h := rgbToHSL(c)
		h.H = rotateHue(h.H, delta)
		out.colors[i] = hslToRGB(h)
	}

	_ = fmt.Sprintf // keep import if needed later
	return out
}
