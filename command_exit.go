package main

import "os"

func commandExit(cfg *config, parameter string) error {
	os.Exit(0)
	return nil
}
