package main

import (
	"fmt"

	"github.com/benedekpal/pokedex/internal/pokeapi"
)

// Handles both 'next' and 'previous' navigation
func handleMapNavigation(config *PokedexConfig, direction string) error {
	// Default URL if no previous/next exists
	url := pokeapi.BaseURL + "/location-area"

	switch direction {
	case "next":
		if config.nextLocationsURL != nil {
			url = *config.nextLocationsURL
		}
	case "prev":
		if config.prevLocationsURL != nil {
			url = *config.prevLocationsURL
		}
	default:
		return fmt.Errorf("invalid direction: %s", direction)
	}

	// Fetch or retrieve from cache
	var rawPageBody []byte
	if cachedBody, found := config.pokeapiCache.Get(url); found {
		rawPageBody = cachedBody
	} else {
		body, err := config.pokeapiClient.GetResponse(&url)
		if err != nil {
			return err
		}
		config.pokeapiCache.Add(url, body)
		rawPageBody = body
	}

	// Decode JSON
	var locationResponse pokeapi.PokeAPILocationResponse
	if err := config.pokeapiClient.DecodeIntoJson(rawPageBody, &locationResponse); err != nil {
		return err
	}

	// Update navigation URLs
	config.nextLocationsURL = locationResponse.Next
	config.prevLocationsURL = locationResponse.Previous

	// Print location names
	for _, result := range locationResponse.Results {
		fmt.Println(result.Name)
	}

	return nil
}

// Next page command
func commandMap(config *PokedexConfig) error {
	return handleMapNavigation(config, "next")
}

// Previous page command
func commandMapb(config *PokedexConfig) error {
	return handleMapNavigation(config, "prev")
}
