package main

import (
	"fmt"
	"log"

	"github.com/arceus-7/chroma16"
)

func main() {
	// Create a palette
	p, err := chroma16.New().
		Seed("matrix").
		Mood(chroma16.Dark).
		Build()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("=== Alacritty ===")
	fmt.Println(p.ToAlacritty())

	fmt.Println("\n=== Kitty ===")
	fmt.Println(p.ToKitty())

	fmt.Println("\n=== Windows Terminal ===")
	fmt.Println(p.ToWindowsTerminal("Matrix Dark"))

	fmt.Println("\n=== Xresources ===")
	fmt.Println(p.ToXresources())
}
