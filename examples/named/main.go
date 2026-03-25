package main

import (
	"fmt"
	"log"

	"github.com/arceus-7/chroma16"
)

func main() {
	// chroma16 supports all 148 standard CSS named colors.
	// We can pass them just like a hex string. They are case-insensitive.
	p, err := chroma16.From("crimson")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=== Generated from named color: crimson ===")
	// Print out the generated hex colors
	for i, h := range p.Hex() {
		fmt.Printf("Slot %2d: %s\n", i, h)
	}

	fmt.Println("\n=== Preview ===")
	p.Preview()
}
