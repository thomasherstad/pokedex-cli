package main

import "fmt"

func explore(cfg *config, parameter string) error {
	encounters, err := cfg.pokeapiClient.GetPokemonFromLocation(parameter)
	if err != nil {
		fmt.Printf("Area not found.\nError: %s\n", err)
		return err
	}
	fmt.Printf("Exploring %s...\n", parameter)
	for _, pokemon := range encounters.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
