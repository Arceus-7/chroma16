
package main

import (
	"fmt"
	"log"

	"github.com/arceus-7/chroma16"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	// Generate a vivid neon theme
	p, err := chroma16.New().
		Seed("cyberpunk").
		Mood(chroma16.Neon).
		Build()
	if err != nil {
		log.Fatal(err)
	}

	// Calling ToLipglossTheme() provides a struct where fields cleanly
	// map to lipgloss.Color values.
	theme := p.ToLipglossTheme()

	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(theme.BrightCyan).
		Background(theme.Black).
		Padding(1, 4).
		MarginBottom(1)

	successStyle := lipgloss.NewStyle().
		Foreground(theme.BrightGreen).
		MarginLeft(2)

	warningStyle := lipgloss.NewStyle().
		Foreground(theme.BrightYellow).
		MarginLeft(2)

	errorStyle := lipgloss.NewStyle().
		Foreground(theme.BrightWhite).
		Background(theme.Red).
		Bold(true).
		Padding(0, 1).
		MarginLeft(2)

	fmt.Println(headerStyle.Render("SYSTEM STATUS DEPLOYMENT REPORT"))
	fmt.Println(successStyle.Render("✔  Core systems online"))
	fmt.Println(warningStyle.Render("⚠  Reactor temperature nominal but rising"))
	fmt.Println(errorStyle.Render("✖  MAINFRAME BREACH DETECTED"))
	fmt.Println()
}
