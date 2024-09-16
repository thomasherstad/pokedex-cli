package main

import (
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {

	locations, err := cfg.pokeapiClient.GetLocations(cfg.nextLocation)
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

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousLocation == nil {
		fmt.Println("You are on the first page")
		return nil
	}

	locations, err := cfg.pokeapiClient.GetLocations(cfg.previousLocation)
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
