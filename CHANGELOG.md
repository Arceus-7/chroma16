# Changelog

All notable changes to this project will be documented in this file.

## [0.1.0] — 2026-03-24

### Added

- `From(seed)` one-liner entry point — accepts `#RRGGBB`, bare `RRGGBB`, or any string
- Builder API: `New().Seed().Mood().Contrast().Build()`
- Six `Mood` constants: `Neutral`, `Warm`, `Cool`, `Dark`, `Pastel`, `Neon`
- Three `Contrast` constants: `Medium`, `High`, `Low`
- `Palette` output methods: `.Hex()`, `.ANSI()`, `.RGB()`, `.At()`, `.Preview()`
- String seed hashing via FNV-1a (zero external dependencies)
- ANSI true-color terminal swatch renderer (Preview)
- Full test suite with ≥80% coverage
