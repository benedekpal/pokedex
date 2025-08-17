package main

import (
	"errors"
	"fmt"
)

func commandHelp(_ *PokedexConfig, _ []string) error {
	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	availableCommands := getCommands()

	if len(availableCommands) == 0 {
		return errors.New("no commands are available")
	}

	for name, cmd := range availableCommands {
		fmt.Printf("%s: %s\n", name, cmd.description)
	}

	return nil
}
