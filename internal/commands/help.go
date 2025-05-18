package commands

import (
	"fmt"

	"github.com/muhola/goPokedex/internal/cache"
	"github.com/muhola/goPokedex/internal/pokeapi"
)

func commandHelp(config *Configuration, pokeCache *cache.Cache, pokeClient *pokeapi.Client, parameter string, pokedex map[string]pokeapi.Pokemon) error {
	commands := GetCommands()
	fmt.Printf("\nWelcome to the Pokedex!\n")
	fmt.Printf("Usage: \n\n")

	for _, value := range commands {
		fmt.Println(value.name, ": ", value.description)
	}
	fmt.Printf("\n")
	return nil
}
