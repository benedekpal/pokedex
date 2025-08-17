package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/benedekpal/pokedex/internal/pokeapi"
	"github.com/benedekpal/pokedex/internal/pokecache"
)

type PokedexConfig struct {
	pokeapiClient    pokeapi.Client
	pokeapiCache     *pokecache.Cache
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *PokedexConfig) {

	availableCommands := getCommands()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		var args []string

		if len(words) > 1 {
			args = words[1:]
		}

		fmt.Printf("Your command was: %s\n", commandName)

		command, exists := availableCommands[commandName]

		if exists {
			err := command.callback(cfg, args)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*PokedexConfig, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the names of locations the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Display the pokemons in the area",
			callback:    commandExplore,
		},
	}
}
