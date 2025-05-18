package commands

import (
	"fmt"

	"github.com/muhola/goPokedex/internal/cache"
	"github.com/muhola/goPokedex/internal/pokeapi"
)

func commandExplore(configuration *Configuration, pokeCache *cache.Cache, pokeClient *pokeapi.Client, pokemon string, pokedex map[string]pokeapi.Pokemon) error {

	pokemonEncounters := pokeClient.GetPokemonBasedOnLocation(pokemon, pokeCache)

	for _, pokemon := range pokemonEncounters.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
