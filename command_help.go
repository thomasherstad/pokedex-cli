package main

import "fmt"

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the pokedex")
	fmt.Println("Type the following commands to get the expected result")
	for _, command := range makeCommands() {
		fmt.Printf("- %v: \n \t %v\n", command.name, command.description)
	}
	return nil
}
