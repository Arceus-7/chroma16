package chroma16

import (
	"fmt"
	"strings"
)

// ToAlacritty returns the palette as an Alacritty TOML colors block.
// Paste this into your alacritty.toml (or import it as a theme file).
//
// Background uses slot 0 (Black), foreground uses slot 15 (Bright White).
func (p Palette) ToAlacritty() string {
	c := p.colors
	// Convert to 0x-prefixed hex for TOML (Alacritty convention)
	hex := func(r rgb) string {
		return fmt.Sprintf("\"#%02X%02X%02X\"", r.R, r.G, r.B)
	}
	var sb strings.Builder
	sb.WriteString("[colors.primary]\n")
	sb.WriteString(fmt.Sprintf("background = %s\n", hex(c[0])))
	sb.WriteString(fmt.Sprintf("foreground = %s\n", hex(c[15])))
	sb.WriteString("\n[colors.normal]\n")
	sb.WriteString(fmt.Sprintf("black   = %s\n", hex(c[0])))
	sb.WriteString(fmt.Sprintf("red     = %s\n", hex(c[1])))
	sb.WriteString(fmt.Sprintf("green   = %s\n", hex(c[2])))
	sb.WriteString(fmt.Sprintf("yellow  = %s\n", hex(c[3])))
	sb.WriteString(fmt.Sprintf("blue    = %s\n", hex(c[4])))
	sb.WriteString(fmt.Sprintf("magenta = %s\n", hex(c[5])))
	sb.WriteString(fmt.Sprintf("cyan    = %s\n", hex(c[6])))
	sb.WriteString(fmt.Sprintf("white   = %s\n", hex(c[7])))
	sb.WriteString("\n[colors.bright]\n")
	sb.WriteString(fmt.Sprintf("black   = %s\n", hex(c[8])))
	sb.WriteString(fmt.Sprintf("red     = %s\n", hex(c[9])))
	sb.WriteString(fmt.Sprintf("green   = %s\n", hex(c[10])))
	sb.WriteString(fmt.Sprintf("yellow  = %s\n", hex(c[11])))
	sb.WriteString(fmt.Sprintf("blue    = %s\n", hex(c[12])))
	sb.WriteString(fmt.Sprintf("magenta = %s\n", hex(c[13])))
	sb.WriteString(fmt.Sprintf("cyan    = %s\n", hex(c[14])))
	sb.WriteString(fmt.Sprintf("white   = %s\n", hex(c[15])))
	return sb.String()
}

// ToKitty returns the palette as a Kitty terminal theme file block.
// Paste this into your kitty.conf or save it as a .conf theme file.
//
// Background uses slot 0, foreground uses slot 15.
func (p Palette) ToKitty() string {
	c := p.colors
	var sb strings.Builder
	for i, col := range c {
		sb.WriteString(fmt.Sprintf("color%-2d  #%02X%02X%02X\n", i, col.R, col.G, col.B))
	}
	sb.WriteString(fmt.Sprintf("\nbackground  #%02X%02X%02X\n", c[0].R, c[0].G, c[0].B))
	sb.WriteString(fmt.Sprintf("foreground  #%02X%02X%02X\n", c[15].R, c[15].G, c[15].B))
	return sb.String()
}

// ToWindowsTerminal returns the palette as a JSON color scheme block
// compatible with Windows Terminal's settings.json "schemes" array.
//
// The name parameter sets the "name" field in the JSON — use the same name
// when referencing the scheme in a profile's "colorScheme" field.
func (p Palette) ToWindowsTerminal(name string) string {
	c := p.colors
	hex := func(r rgb) string {
		return fmt.Sprintf("#%02X%02X%02X", r.R, r.G, r.B)
	}
	var sb strings.Builder
	sb.WriteString("{\n")
	sb.WriteString(fmt.Sprintf("    \"name\": %q,\n", name))
	sb.WriteString(fmt.Sprintf("    \"background\": %q,\n", hex(c[0])))
	sb.WriteString(fmt.Sprintf("    \"foreground\": %q,\n", hex(c[15])))
	sb.WriteString(fmt.Sprintf("    \"cursorColor\": %q,\n", hex(c[15])))
	sb.WriteString(fmt.Sprintf("    \"selectionBackground\": %q,\n", hex(c[8])))
	sb.WriteString(fmt.Sprintf("    \"black\": %q,\n", hex(c[0])))
	sb.WriteString(fmt.Sprintf("    \"red\": %q,\n", hex(c[1])))
	sb.WriteString(fmt.Sprintf("    \"green\": %q,\n", hex(c[2])))
	sb.WriteString(fmt.Sprintf("    \"yellow\": %q,\n", hex(c[3])))
	sb.WriteString(fmt.Sprintf("    \"blue\": %q,\n", hex(c[4])))
	sb.WriteString(fmt.Sprintf("    \"purple\": %q,\n", hex(c[5])))
	sb.WriteString(fmt.Sprintf("    \"cyan\": %q,\n", hex(c[6])))
	sb.WriteString(fmt.Sprintf("    \"white\": %q,\n", hex(c[7])))
	sb.WriteString(fmt.Sprintf("    \"brightBlack\": %q,\n", hex(c[8])))
	sb.WriteString(fmt.Sprintf("    \"brightRed\": %q,\n", hex(c[9])))
	sb.WriteString(fmt.Sprintf("    \"brightGreen\": %q,\n", hex(c[10])))
	sb.WriteString(fmt.Sprintf("    \"brightYellow\": %q,\n", hex(c[11])))
	sb.WriteString(fmt.Sprintf("    \"brightBlue\": %q,\n", hex(c[12])))
	sb.WriteString(fmt.Sprintf("    \"brightPurple\": %q,\n", hex(c[13])))
	sb.WriteString(fmt.Sprintf("    \"brightCyan\": %q,\n", hex(c[14])))
	sb.WriteString(fmt.Sprintf("    \"brightWhite\": %q\n", hex(c[15])))
	sb.WriteString("}")
	return sb.String()
}

// ToXresources returns the palette as an X resources color block.
// Paste this into your ~/.Xresources file and run `xrdb -merge ~/.Xresources`.
//
// Background uses slot 0, foreground uses slot 15.
func (p Palette) ToXresources() string {
	c := p.colors
	var sb strings.Builder
	for i, col := range c {
		sb.WriteString(fmt.Sprintf("*.color%-2d:  #%02X%02X%02X\n", i, col.R, col.G, col.B))
	}
	sb.WriteString(fmt.Sprintf("\n*.background: #%02X%02X%02X\n", c[0].R, c[0].G, c[0].B))
	sb.WriteString(fmt.Sprintf("*.foreground: #%02X%02X%02X\n", c[15].R, c[15].G, c[15].B))
	return sb.String()
}
