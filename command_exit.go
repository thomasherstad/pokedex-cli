package main

import "os"


func commandExit(commands map[string]cliCommand) error {
	os.Exit(0)
	return nil
}