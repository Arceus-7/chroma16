package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/arceus-7/chroma16"
)

func main() {
	// 1. Generate a palette
	orig, err := chroma16.From("#123456")
	if err != nil {
		log.Fatal(err)
	}

	// 2. Marshal to JSON
	// The Palette struct implements json.Marshaler natively.
	data, err := json.MarshalIndent(orig, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Serialized JSON:\n%s\n\n", string(data))

	// 3. Unmarshal back to a Palette
	// chroma16 provides a helper function to easily restore from bytes
	restored, err := chroma16.FromJSON(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Restored Palette Matches Original?")
	origHex, _ := orig.At(0)
	restoredHex, _ := restored.At(0)
	fmt.Printf("Slot 0 -> Orig: %s | Restored: %s\n", origHex, restoredHex)
}
