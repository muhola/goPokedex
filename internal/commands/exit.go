package commands

import (
	"fmt"
	"os"

	"github.com/muhola/goPokedex/internal/cache"
	"github.com/muhola/goPokedex/internal/pokeapi"
)

func commandExit(config *Configuration, pokeCache *cache.Cache, pokeClient *pokeapi.Client, parameter string, pokedex map[string]pokeapi.Pokemon) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return nil
}
