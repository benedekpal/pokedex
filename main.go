package main

import (
	"time"

	"github.com/benedekpal/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &PokedexConfig{
		pokeapiClient:    pokeClient,
		nextLocationsURL: nil,
		prevLocationsURL: nil,
	}
	startRepl(cfg)
}
