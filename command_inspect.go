package main

import (
	"fmt"
	"pokedex-cli/internal/pokeapi"
)

func inspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		fmt.Println("you can only inspect one pokemon at a time")
		return nil
	}

	if _, ok := cfg.caughtPokemon[args[0]]; !ok {
		fmt.Printf("You have not caught %s\n", args[0])
		return nil
	}
	printInspection(cfg.caughtPokemon[args[0]])
	return nil
}

func printInspection(p pokeapi.Pokemon) {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %v\n", p.Height)
	fmt.Printf("Weight: %v\n", p.Weight)
	fmt.Println("Stats: ")
	//Stats loop
	for _, stat := range p.Stats {
		fmt.Printf("	- %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types: ")
	for _, t := range p.Types {
		fmt.Printf("	- %s\n", t.Type.Name)
	}
}
