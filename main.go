package main

import (
	"time"

	"github.com/benedekpal/pokedex/internal/pokeapi"
	"github.com/benedekpal/pokedex/internal/pokecache"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(300 * time.Second)
	cfg := &PokedexConfig{
		pokeapiCache:     pokeCache,
		pokeapiClient:    pokeClient,
		nextLocationsURL: nil,
		prevLocationsURL: nil,
		pokedex: &pokeapi.Pokedex{
			Pokemons: make(map[string]pokeapi.Pokemon),
		},
	}
	startRepl(cfg)
}
