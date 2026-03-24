# Changelog

All notable changes to this project will be documented in this file.

## [0.2.0] — 2026-03-24

### Added

- **CSS Named Colors**: Support for all 148 CSS named colors (e.g., `From("crimson")`).
- **Export Formats**: Added `.ToAlacritty()`, `.ToKitty()`, `.ToWindowsTerminal("Name")`, and `.ToXresources()` to generate config blocks.
- **Palette Manipulation**: Added `Blend(p1, p2, t)`, `.Complement()`, and `.Analogous(degrees)` for dynamic palette adjustment.
- **JSON Serialization**: Added `.MarshalJSON()`, `.UnmarshalJSON()`, and `FromJSON(data)` for saving/restoring palettes.
- **Lipgloss Integration**: Added `.ToLipglossTheme()` (requires building with `-tags lipgloss`).

## [0.1.2] — 2026-03-24

### Fixed
- Changed module path to lowercase `github.com/arceus-7/chroma16` for Go proxy compatibility.

### Added
- MIT License
- Three runnable example programs in `examples/`

## [0.1.0] — 2026-03-24

### Added
- Initial release
