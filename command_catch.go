package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/benedekpal/pokedex/internal/pokeapi"
)

func commandCatch(config *PokedexConfig, args []string) error {
	if len(args) != 1 {
		return errors.New("missing argument")
	}
	if args[0] == "" {
		return errors.New("pokemon name is missing")
	}

	pokemon, err := getPokemonStats(config, args[0])

	if err != nil {
		return err
	} else {
		fmt.Println("Throwing a Pokeball at " + args[0] + "...")
		if attempCatch(pokemon.BaseExperiance) {
			fmt.Println(args[0] + " was caught!")
			config.pokedex.Pokemons[args[0]] = pokemon
		} else {
			fmt.Println(args[0] + " escaped!")
		}
	}

	return nil
}

// separated and reworked function so it can be tested as requested
func getPokemonStats(config *PokedexConfig, pokemonName string) (pokeapi.Pokemon, error) {

	var pokemonData pokeapi.Pokemon
	url := pokeapi.BaseURL + "/pokemon/" + pokemonName

	// Fetch or retrieve from cache
	var rawPageBody []byte
	if cachedBody, found := config.pokeapiCache.Get(url); found {
		rawPageBody = cachedBody
	} else {
		body, err := config.pokeapiClient.GetResponse(&url)
		if err != nil {
			return pokemonData, err
		}
		config.pokeapiCache.Add(url, body)
		rawPageBody = body
	}

	if err := config.pokeapiClient.DecodeIntoJson(rawPageBody, &pokemonData); err != nil {
		return pokemonData, err
	}

	return pokemonData, nil
}

func attempCatch(pokemonExp int) bool {
	randomNum := rand.Intn(pokeapi.MaxExp-pokeapi.MinExp) + pokeapi.MinExp
	if randomNum < pokemonExp {
		return false
	}
	return true
}
