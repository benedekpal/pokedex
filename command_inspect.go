package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/benedekpal/pokedex/internal/pokeapi"
)

func commandInspect(config *PokedexConfig, args []string) error {
	if len(args) != 1 {
		return errors.New("missing argument")
	}
	if args[0] == "" {
		return errors.New("pokemon name is missing")
	}

	value, ok := config.pokedex.Pokemons[args[0]]
	if ok {
		fmt.Println(PrintPretty(value))
	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}

// Pretty print method
func PrintPretty(p pokeapi.Pokemon) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Name: %s\n", p.Name))
	sb.WriteString(fmt.Sprintf("Height: %d\n", p.Height))
	sb.WriteString(fmt.Sprintf("Weight: %d\n", p.Weight))

	sb.WriteString("Stats:\n")
	for _, stat := range p.Stats {
		sb.WriteString(fmt.Sprintf("  -%s: %d\n", stat.Stat.StatName, stat.Base_Stat))
	}

	sb.WriteString("Types:\n")
	for i, t := range p.Types {
		if i == len(p.Types)-1 {
			sb.WriteString(fmt.Sprintf("  - %s", t.Type.Name)) // no newline
		} else {
			sb.WriteString(fmt.Sprintf("  - %s\n", t.Type.Name))
		}
	}

	return sb.String()
}
