package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore which Pokemon are in the chosen area",
			callback:    explore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a pokemon of your choice",
			callback:    catch,
		},
		"inspect": {
			name:        "inspect",
			description: "See information about a pokemon you have caught",
			callback:    inspect,
		},
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	return strings.Fields(lower)
}

func startRepl(cfg *config) {
	commands := makeCommands()
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("pokedex > ")
		scanner.Scan()
		text := scanner.Text()

		if len(text) == 0 {
			continue
		}

		words := cleanInput(text)

		commandName := words[0]
		var args []string
		if len(words) > 1 {
			args = words[1:]
		}

		command, ok := commands[commandName]
		if ok {
			command.callback(cfg, args...)
		}
	}
}
