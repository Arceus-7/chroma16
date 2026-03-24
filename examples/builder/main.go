// builder demonstrates the fluent Builder API with every Mood and Contrast combo.
// Run: go run ./examples/builder
package main

import (
	"fmt"
	"log"

	"github.com/arceus-7/chroma16"
)

func main() {
	seed := "#9B5DE5" // vibrant purple

	type combo struct {
		mood     chroma16.Mood
		moodName string
		contrast chroma16.Contrast
		contName string
	}

	combos := []combo{
		{chroma16.Neutral, "Neutral", chroma16.Medium, "Medium"},
		{chroma16.Warm, "Warm", chroma16.High, "High"},
		{chroma16.Cool, "Cool", chroma16.High, "High"},
		{chroma16.Dark, "Dark", chroma16.Medium, "Medium"},
		{chroma16.Pastel, "Pastel", chroma16.Low, "Low"},
		{chroma16.Neon, "Neon", chroma16.High, "High"},
	}

	for _, c := range combos {
		p, err := chroma16.New().
			Seed(seed).
			Mood(c.mood).
			Contrast(c.contrast).
			Build()
		if err != nil {
			log.Fatalf("build error: %v", err)
		}

		fmt.Printf("── Mood: %-8s  Contrast: %-6s  Seed: %s ──\n",
			c.moodName, c.contName, seed)
		p.Preview()
	}
}
