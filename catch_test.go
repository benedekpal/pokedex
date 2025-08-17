package main

import (
	"testing"
	"time"

	"github.com/benedekpal/pokedex/internal/pokeapi"
	"github.com/benedekpal/pokedex/internal/pokecache"
)

func TestGetPokemonStats(t *testing.T) {
	cases := []struct {
		input    string
		expected pokeapi.Pokemon
	}{
		{
			input: "pikachu",
			expected: pokeapi.Pokemon{
				Name:           "pikachu",
				BaseExperiance: 112,
			},
		},
	}

	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(5 * time.Second)
	cfg := &PokedexConfig{
		pokeapiCache:     pokeCache,
		pokeapiClient:    pokeClient,
		nextLocationsURL: nil,
		prevLocationsURL: nil,
		pokedex: &pokeapi.Pokedex{
			Pokemons: make(map[string]pokeapi.Pokemon),
		},
	}

	for _, c := range cases {
		actual, err := getPokemonStats(cfg, c.input)
		if err != nil {
			t.Fatalf("getPokemonStats failed: %v", err)
		}

		if actual.BaseExperiance != c.expected.BaseExperiance {
			t.Fatalf("values do not match: got %d, expected %d", actual.BaseExperiance, c.expected.BaseExperiance)
		}
	}
}
