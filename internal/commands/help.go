package commands

import (
	"fmt"

	"github.com/muhola/goPokedex/internal/cache"
	"github.com/muhola/goPokedex/internal/pokeapi"
)

func commandHelp(config *Configuration, pokeCache *cache.Cache, pokeClient *pokeapi.Client) error {
	commands := GetCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage: \n\n")

	for _, value := range commands {
		fmt.Println(value.name, ": ", value.description)
	}
	fmt.Printf("\n")
	return nil
}
