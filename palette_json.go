package chroma16

import (
	"encoding/json"
	"fmt"
)

// paletteMarshalProxy is the on-wire JSON shape for a Palette.
type paletteMarshalProxy struct {
	Version int      `json:"version"`
	Colors  []string `json:"colors"`
}

// MarshalJSON implements encoding/json.Marshaler.
// The JSON representation contains a version field and an array of 16
// "#RRGGBB" hex strings ordered by terminal slot index (0–15).
func (p Palette) MarshalJSON() ([]byte, error) {
	return json.Marshal(paletteMarshalProxy{
		Version: 1,
		Colors:  p.Hex(),
	})
}

// UnmarshalJSON implements encoding/json.Unmarshaler.
// It expects the JSON format produced by MarshalJSON (version + 16-element
// colors array). The version field is checked but only version 1 is currently
// supported.
func (p *Palette) UnmarshalJSON(data []byte) error {
	var proxy paletteMarshalProxy
	if err := json.Unmarshal(data, &proxy); err != nil {
		return fmt.Errorf("chroma16: invalid palette JSON: %w", err)
	}
	if proxy.Version != 1 {
		return fmt.Errorf("chroma16: unsupported palette JSON version %d (want 1)", proxy.Version)
	}
	if len(proxy.Colors) != 16 {
		return fmt.Errorf("chroma16: palette JSON must have exactly 16 colors, got %d", len(proxy.Colors))
	}
	for i, h := range proxy.Colors {
		c, err := parseHex(h)
		if err != nil {
			return fmt.Errorf("chroma16: invalid color at index %d: %w", i, err)
		}
		p.colors[i] = c
	}
	return nil
}

// FromJSON reconstructs a Palette from JSON produced by Palette.MarshalJSON.
// Returns an error if the data is malformed, missing colors, or contains
// invalid hex values.
func FromJSON(data []byte) (Palette, error) {
	var p Palette
	if err := json.Unmarshal(data, &p); err != nil {
		return Palette{}, err
	}
	return p, nil
}
