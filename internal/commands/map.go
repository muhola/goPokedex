package commands

import (
	"fmt"

	"github.com/muhola/goPokedex/internal/cache"
	"github.com/muhola/goPokedex/internal/pokeapi"
)

func commandMap(configuration *Configuration, pokeCache *cache.Cache, pokeClient *pokeapi.Client, parameter string, pokedex map[string]pokeapi.Pokemon) error {
	url := ""

	if configuration.Next != "" {
		url = configuration.Next
	}
	locations := pokeClient.GetLocation(url, pokeCache)

	if locations.Next != "<nil>" {
		configuration.Next = locations.Next
	}
	if locations.Previous != "<nil>" {
		configuration.Previous = locations.Previous
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}
