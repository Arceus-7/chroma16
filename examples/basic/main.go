// basic demonstrates the one-liner From() API.
// Run: go run ./examples/basic
package main

import (
	"fmt"
	"log"

	"github.com/arceus-7/chroma16"
)

func main() {
	seeds := []string{
		"#FF6B35", // warm orange
		"#2E86AB", // cool blue
		"ocean",   // string seed — hashed deterministically
		"forest",
		"sunset",
	}

	for _, seed := range seeds {
		p, err := chroma16.From(seed)
		if err != nil {
			log.Fatalf("error building palette for %q: %v", seed, err)
		}

		fmt.Printf("── Seed: %q ──\n", seed)
		p.Preview()
	}
}
