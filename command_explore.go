package main

import (
	"errors"
	"fmt"
)

func explore(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("no area given")
	} else if len(args) > 1 {
		return errors.New("too many areas given")
	}

	loc := args[0]
	encounters, err := cfg.pokeapiClient.GetPokemonFromLocation(loc)
	if err != nil {
		fmt.Printf("Area not found.\nError: %s\n", err)
		return err
	}
	fmt.Printf("Exploring %s...\n", loc)
	for _, pokemon := range encounters.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
