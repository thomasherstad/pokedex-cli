package main

import (
	"errors"
	"fmt"
)

func pokedex(cfg *config, args ...string) error {
	if len(args) != 0 {
		return errors.New("too many arguments")
	}

	fmt.Println("---POKEDEX---")
	for _, p := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", p.Name)
	}
	return nil
}
