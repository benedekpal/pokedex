package main

import (
	"fmt"

	"github.com/benedekpal/pokedex/internal/pokeapi"
)

func commandMap(config *PokedexConfig) error {
	// If there's no next page, return
	url := pokeapi.BaseURL + "/location-area"
	if config.nextLocationsURL != nil {
		url = *config.nextLocationsURL
	}

	var locationResponse PokeAPILocationResponse

	err := config.pokeapiClient.GetJsonResponseAndDecode(&url, &locationResponse)
	if err != nil {
		return err
	}

	config.nextLocationsURL = locationResponse.Next
	config.prevLocationsURL = locationResponse.Previous

	// Print the names of locations
	for _, result := range locationResponse.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapb(config *PokedexConfig) error {
	// If there's no next page, return
	url := pokeapi.BaseURL + "/location-area"
	if config.nextLocationsURL != nil {
		url = *config.nextLocationsURL
	}

	var locationResponse PokeAPILocationResponse

	err := config.pokeapiClient.GetJsonResponseAndDecode(&url, &locationResponse)
	if err != nil {
		return err
	}

	config.nextLocationsURL = locationResponse.Next
	config.prevLocationsURL = locationResponse.Previous

	// Print the names of locations
	for _, result := range locationResponse.Results {
		fmt.Println(result.Name)
	}

	return nil
}
