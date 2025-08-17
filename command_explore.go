package main

import (
	"errors"
	"fmt"

	"github.com/benedekpal/pokedex/internal/pokeapi"
)

func commandExplore(config *PokedexConfig, args []string) error {
	if len(args) != 1 {
		return errors.New("missing argument")
	}
	if args[0] == "" {
		return errors.New("area name is missing")
	}

	pokemons, err := handleExploration(config, args[0])

	if err != nil {
		return err
	} else {
		fmt.Println("Exploring " + args[0] + "...")
		fmt.Println("Found Pokemon:")
		// Print pokemon names
		for _, result := range pokemons.PokemonEncounters {
			fmt.Println(" - " + result.Pokemon.Name)
		}
	}

	return nil
}

// separated and reworked function so it can be tested as requested
func handleExploration(config *PokedexConfig, areaName string) (pokeapi.MinimalLocationArea, error) {

	var locationData pokeapi.MinimalLocationArea
	url := pokeapi.BaseURL + "/location-area/" + areaName

	// Fetch or retrieve from cache
	var rawPageBody []byte
	if cachedBody, found := config.pokeapiCache.Get(url); found {
		rawPageBody = cachedBody
	} else {
		body, err := config.pokeapiClient.GetResponse(&url)
		if err != nil {
			return locationData, err
		}
		config.pokeapiCache.Add(url, body)
		rawPageBody = body
	}

	if err := config.pokeapiClient.DecodeIntoJson(rawPageBody, &locationData); err != nil {
		return locationData, err
	}

	return locationData, nil
}
