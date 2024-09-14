package main

import (
	"fmt"
	"pokedex-cli/internal/pokeapi"
)

func commandMapf(cfg *config) error {
	locations, err := pokeapi.GetLocations(cfg.nextLocation)
	if err != nil {
		return err
	}
	cfg.nextLocation = locations.Next
	cfg.previousLocation = locations.Previous

	for _, loc := range locations.Locations {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previousLocation == nil {
		fmt.Println("You are on the first page")
		return nil
	}

	locations, err := pokeapi.GetLocations(cfg.previousLocation)
	if err != nil {
		return err
	}
	cfg.nextLocation = locations.Next
	cfg.previousLocation = locations.Previous

	for _, loc := range locations.Locations {
		fmt.Println(loc.Name)
	}
	return nil
}
