package commands

import (
	"fmt"

	"github.com/muhola/goPokedex/internal/cache"
	"github.com/muhola/goPokedex/internal/pokeapi"
)

func commandPokedex(config *Configuration, pokeCache *cache.Cache, pokeClient *pokeapi.Client, parameter string, pokedex map[string]pokeapi.Pokemon) error {

	if len(pokedex) == 0 {
		fmt.Println("Your Pokedex is empty, try catch a pokemon with the 'catch' command ")
		return nil
	}

	for _, pokemon := range pokedex {
		fmt.Println(pokemon.Name)
	}

	return nil
}
