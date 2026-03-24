package chroma16

// Mood controls the overall temperature and saturation feel of the palette.
type Mood int

const (
	// Neutral is the default mood — balanced temperature, no shift applied.
	Neutral Mood = iota
	// Warm shifts the palette toward reds and oranges.
	Warm
	// Cool shifts the palette toward blues and cyans.
	Cool
	// Dark produces lower brightness variants throughout.
	Dark
	// Pastel desaturates the palette for soft, muted tones.
	Pastel
	// Neon cranks saturation for vivid, high-energy colors.
	Neon
)

// Contrast controls the luminance spread across the 16 color slots.
type Contrast int

const (
	// Medium is the default — balanced luminance range.
	Medium Contrast = iota
	// High widens the luminance range for stronger contrast.
	High
	// Low narrows the range for a softer, more uniform look.
	Low
)

// moodAdjust holds per-Mood HSL adjustment parameters used by generateColors.
type moodAdjust struct {
	HueShift float64 // degrees to rotate all hues
	SatMul   float64 // multiply saturation by this factor
	LightOff float64 // add this offset to lightness (can be negative)
}

// moodParams returns the HSL adjustment parameters for the given Mood.
func moodParams(m Mood) moodAdjust {
	switch m {
	case Warm:
		return moodAdjust{HueShift: -15, SatMul: 1.1, LightOff: 0}
	case Cool:
		return moodAdjust{HueShift: 30, SatMul: 0.95, LightOff: 0.02}
	case Dark:
		return moodAdjust{HueShift: 0, SatMul: 0.85, LightOff: -0.08}
	case Pastel:
		return moodAdjust{HueShift: 0, SatMul: 0.45, LightOff: 0.15}
	case Neon:
		return moodAdjust{HueShift: 0, SatMul: 1.4, LightOff: 0}
	default: // Neutral
		return moodAdjust{HueShift: 0, SatMul: 1.0, LightOff: 0}
	}
}

// contrastSpread returns the lightness anchors for normal (slots 0–7) and
// bright (slots 8–15) color groups based on the requested Contrast level.
func contrastSpread(c Contrast) (float64, float64) {
	switch c {
	case High:
		return 0.30, 0.72
	case Low:
		return 0.42, 0.60
	default: // Medium
		return 0.36, 0.66
	}
}
