package main

import "pokedex-cli/internal/pokeapi"

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocation     *string
	previousLocation *string
}
