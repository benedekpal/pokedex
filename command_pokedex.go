package main

import (
	"fmt"
)

func commandPokedex(config *PokedexConfig, _ []string) error {
	fmt.Println("Your Pokedex:")

	for name, _ := range config.pokedex.Pokemons {
		fmt.Println(" - " + name)
	}
	return nil
}
