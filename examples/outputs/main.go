// outputs shows every output format: Hex, RGB, ANSI, and At().
// Run: go run ./examples/outputs
package main

import (
	"fmt"
	"log"

	"github.com/arceus-7/chroma16"
)

func main() {
	p, err := chroma16.From("#00B4D8")
	if err != nil {
		log.Fatal(err)
	}

	// ── Hex strings ─────────────────────────────────────────────────────────
	fmt.Println("=== Hex() ===")
	for i, h := range p.Hex() {
		fmt.Printf("  [%2d] %s\n", i, h)
	}

	// ── Raw RGB triples ──────────────────────────────────────────────────────
	fmt.Println("\n=== RGB() ===")
	for i, rgb := range p.RGB() {
		fmt.Printf("  [%2d] R=%-3d G=%-3d B=%-3d\n", i, rgb[0], rgb[1], rgb[2])
	}

	// ── ANSI OSC 4 sequences ─────────────────────────────────────────────────
	fmt.Println("\n=== ANSI() ===")
	fmt.Println("  (sequences printed below — paste into a supported terminal to repaint it)")
	for _, seq := range p.ANSI() {
		// Print the raw bytes so they're visible in stdout
		fmt.Printf("  %q\n", seq)
	}

	// ── At() single index ────────────────────────────────────────────────────
	fmt.Println("\n=== At() ===")
	for _, idx := range []int{0, 7, 8, 15} {
		color, err := p.At(idx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("  slot %2d → %s\n", idx, color)
	}

	// ── Visual swatch ────────────────────────────────────────────────────────
	fmt.Println("\n=== Preview() ===")
	p.Preview()
}
