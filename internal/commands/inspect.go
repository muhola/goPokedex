package commands

import (
	"fmt"
	"strings"

	"github.com/muhola/goPokedex/internal/cache"
	"github.com/muhola/goPokedex/internal/pokeapi"
)

func commandIncpect(config *Configuration, pokeCache *cache.Cache, pokeClient *pokeapi.Client, parameter string, pokedex map[string]pokeapi.Pokemon) error {
	if len(pokedex) == 0 {
		fmt.Println("Your Pokedex is empty, try catch a pokemon with the 'catch' command ")
		return nil
	}
	if strings.TrimSpace(parameter) == "" {
		fmt.Printf("Please provied pokemon name, inspect <pokemon name> \n")
		return nil
	}

	getPokemon, exists := pokedex[parameter]
	if !exists {
		fmt.Printf("Can't find any pokemon called %s\n", parameter)
		return nil
	}

	fmt.Println(buildPokemonOutput(getPokemon))
	return nil
}

func buildPokemonOutput(pokemon pokeapi.Pokemon) string {
	var builder strings.Builder
	builder.WriteByte('\n')

	builder.WriteString(fmt.Sprintf("Name: %s\n", pokemon.Name))

	builder.WriteString(fmt.Sprintf("Weight: %d\n", pokemon.Weight))
	builder.WriteString(fmt.Sprintf("Height: %d\n", pokemon.Height))

	builder.WriteByte('\n')

	builder.WriteString("Stats:\n")
	for _, stats := range pokemon.Stats {
		builder.WriteString(fmt.Sprintf("  - %s: %d\n", stats.Stat.Name, stats.BaseStat))
	}
	builder.WriteString("Types:\n")
	for _, pokemonType := range pokemon.Types {
		builder.WriteString(fmt.Sprintf("  - %s \n", pokemonType.Type.Name))
	}

	// Get the final string
	return builder.String()
}
