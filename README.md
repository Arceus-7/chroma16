# chroma16
[![codecov](https://codecov.io/gh/arceus-7/chroma16/graph/badge.svg)](https://app.codecov.io/github/Arceus-7/chroma16)
[![Go Reference](https://pkg.go.dev/badge/github.com/arceus-7/chroma16.svg)](https://pkg.go.dev/github.com/arceus-7/chroma16)
[![Go ≥1.21](https://img.shields.io/badge/go-%3E%3D1.21-blue)](https://go.dev/dl)
[![License: MIT](https://img.shields.io/badge/license-MIT-green)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/arceus-7/chroma16)](https://goreportcard.com/report/github.com/arceus-7/chroma16)

Generate a complete, harmonious 16-color terminal palette from a single seed value — a hex color or any string.

---

## Install

```bash
go get github.com/arceus-7/chroma16
```

---

## Quick Start

```go
import "github.com/arceus-7/chroma16"

palette, err := chroma16.From("#FF6B35")
if err != nil {
    log.Fatal(err)
}
palette.Preview() // prints a color swatch to your terminal
```

---

## Builder API

```go
palette, err := chroma16.New().
    Seed("#2E86AB").
    Mood(chroma16.Cool).
    Contrast(chroma16.High).
    Build()
if err != nil {
    log.Fatal(err)
}

// Get all 16 hex colors
hexColors := palette.Hex()

// Apply to terminal session with ANSI OSC 4 sequences
for _, seq := range palette.ANSI() {
    fmt.Print(seq)
}
```

---

## Examples

Three runnable examples live in the `examples/` directory:

| Program | Run | What it shows |
|---------|-----|---------------|
| `examples/basic` | `go run ./examples/basic` | `From()` with hex and string seeds |
| `examples/builder` | `go run ./examples/builder` | Every `Mood` × `Contrast` combo via the Builder |
| `examples/outputs` | `go run ./examples/outputs` | All output formats: Hex, RGB, ANSI, At(), Preview() |

---



## The 16 Slots

Colors are indexed to match the standard 16-color terminal palette:


| Index | Name         | Index | Name          |
|-------|--------------|-------|---------------|
| 0     | Black        | 8     | Bright Black  |
| 1     | Red          | 9     | Bright Red    |
| 2     | Green        | 10    | Bright Green  |
| 3     | Yellow       | 11    | Bright Yellow |
| 4     | Blue         | 12    | Bright Blue   |
| 5     | Magenta      | 13    | Bright Magenta|
| 6     | Cyan         | 14    | Bright Cyan   |
| 7     | White        | 15    | Bright White  |

---

## Seed Input

| Seed type       | Behaviour                                        |
|-----------------|--------------------------------------------------|
| `"#FF6B35"`     | Parsed as a hex color                            |
| `"FF6B35"`      | Same — leading `#` is optional                   |
| `"ocean"`       | Deterministically hashed via FNV-1a              |
| `""`            | Hashed (FNV-1a of empty string) — valid, stable  |
| `"#GGGGGG"`     | Returns an error (invalid hex characters)        |

> **Note:** Named colors like `"red"` are hashed, not looked up by name.
> Named color lookup is planned for v0.2.

---

## Moods

| Constant         | Effect                                  |
|------------------|-----------------------------------------|
| `chroma16.Neutral` | Default — balanced hues               |
| `chroma16.Warm`    | Shift toward reds and oranges          |
| `chroma16.Cool`    | Shift toward blues and cyans           |
| `chroma16.Dark`    | Lower brightness throughout            |
| `chroma16.Pastel`  | Desaturated, soft, muted tones         |
| `chroma16.Neon`    | High saturation, vivid and energetic   |

---

## Contrast

| Constant          | Effect                              |
|-------------------|-------------------------------------|
| `chroma16.Medium` | Default — balanced luminance range  |
| `chroma16.High`   | Wider luminance spread              |
| `chroma16.Low`    | Narrower, softer luminance range    |

---

## API Reference

### Functions

| Function             | Returns          | Description                                    |
|----------------------|------------------|------------------------------------------------|
| `From(seed string)`  | `Palette, error` | One-liner palette from a seed                  |
| `New()`              | `*Builder`       | Create a Builder for a customized palette      |
| `Blend(p1, p2, t)`   | `Palette`        | Interpolate two palettes (t is 0 to 1)         |
| `FromJSON(data)`     | `Palette, error` | Restore a palette from JSON bytes              |

### Palette Methods

| Method       | Returns        | Description                                            |
|--------------|----------------|--------------------------------------------------------|
| `Hex()`      | `[]string`     | 16 hex strings in `#RRGGBB` format                     |
| `ANSI()`     | `[]string`     | 16 XTerm OSC 4 escape sequences                        |
| `RGB()`      | `[][3]uint8`   | 16 raw `[R, G, B]` triples                             |
| `At(i int)`  | `string, error`| Hex string for slot `i`; error if `i` outside [0,15]  |
| `Preview()`  | —              | Prints a color swatch to stdout                        |
| `Complement()`| `Palette`     | Returns a new palette with opposite hues               |
| `Analogous(d)`| `Palette`     | Returns a new palette with hues shifted by `d` degrees |
| `ToAlacritty()`| `string`     | Returns an Alacritty TOML colors block                 |
| `ToKitty()`   | `string`      | Returns a Kitty `.conf` text block                     |
| `ToWindowsTerminal(n)`| `string` | Returns a JSON block for Windows Terminal          |
| `ToXresources()`| `string`    | Returns an Xresources colors block                     |
| `MarshalJSON()`| `[]byte, err`| Serializes palette to JSON                             |
| `ToLipglossTheme()`| `LipglossTheme` | Available when built with `-tags lipgloss`      |

### Builder Methods

| Method                  | Returns      | Description                         |
|-------------------------|--------------|-------------------------------------|
| `Seed(seed string)`     | `*Builder`   | Set the seed value                  |
| `Mood(m Mood)`          | `*Builder`   | Set the palette mood                |
| `Contrast(c Contrast)`  | `*Builder`   | Set the contrast level              |
| `Build()`               | `Palette, error` | Generate the final palette      |

---

## Color Generation Algorithm

`chroma16` converts the seed to HSL, then places 8 hue relationships
(complement, primary, triadic, analogous) across the 16 slots. Normal colors
(0–7) use lower lightness; bright colors (8–15) use higher lightness. Mood
adjusts the hue shift, saturation multiplier, and lightness offset. Contrast
widens or narrows the lightness spread between the two groups.

All math is pure Go standard library — zero transitive dependencies.

---

## License

MIT © 2026 Arceus-7
