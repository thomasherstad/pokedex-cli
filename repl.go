package main

import (
	"bufio"
	"fmt"
	"os"
)

type command interface {
	Action()
}

type cliCommand struct {
	name        string
	description string
	callback    func(cgf *config) error
}

func makeCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Lists commands and tells you how to use the program",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Close the program",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations in the Pokemon world",
			callback:    commandMapb,
		},
	}
}

func startRepl() {
	commands := makeCommands()

	cfg := config{
		nextLocation:     "https://pokeapi.co/api/v2/location/",
		previousLocation: "",
	}

	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("pokedex > ")
		scanner.Scan()
		text := scanner.Text()

		if len(text) == 0 {
			continue
		}

		if text == "exit" {
			return
		}

		command, ok := commands[text]
		if ok {
			command.callback(&cfg)
		}
	}
}
