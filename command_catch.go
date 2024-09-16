package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func catch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("no pokemon given")
	} else if len(args) > 1 {
		return errors.New("too many pokemon given")
	}

	name := args[0]

	//Can add escape if pokemon has already been caught

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Tried to catch %s\n", pokemon.Name)

	// Chance of catching
	chance := rand.Float32()

	maxBaseExperience := 390
	normalizedExperience := float32(pokemon.BaseExperience) / float32(maxBaseExperience)

	fmt.Println(chance)
	fmt.Println(normalizedExperience)
	if chance >= normalizedExperience {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.caughtPokemon[pokemon.Name] = pokemon
		fmt.Println("Added to pokedex")
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
