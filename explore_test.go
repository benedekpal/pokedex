package main

import (
	"testing"
	"time"

	"github.com/benedekpal/pokedex/internal/pokeapi"
	"github.com/benedekpal/pokedex/internal/pokecache"
)

func TestExplore(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "mt-coronet-2f",
			expected: []string{"clefairy", "golbat", "machoke", "graveler", "nosepass", "medicham", "lunatone", "solrock", "chingling", "bronzor", "bronzong"},
		},
	}

	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokeCache := pokecache.NewCache(5 * time.Second)
	cfg := &PokedexConfig{
		pokeapiCache:     pokeCache,
		pokeapiClient:    pokeClient,
		nextLocationsURL: nil,
		prevLocationsURL: nil,
	}

	for _, c := range cases {
		actual, err := handleExploration(cfg, c.input)
		if err != nil {
			t.Fatalf("handleExploration failed: %v", err)
		}

		if len(actual.PokemonEncounters) != len(c.expected) {
			t.Fatalf("lengths don't match: got %d, expected %d", len(actual.PokemonEncounters), len(c.expected))
		}

		for i, encounter := range actual.PokemonEncounters {
			expectedName := c.expected[i]
			if encounter.Pokemon.Name != expectedName {
				t.Errorf("Pokemon at index %d: got %s, expected %s", i, encounter.Pokemon.Name, expectedName)
			}
		}
	}
}
