package commands

import (
	"fmt"

	"github.com/muhola/goPokedex/internal/cache"
	"github.com/muhola/goPokedex/internal/pokeapi"
)

func commandMapb(configuration *Configuration, pokeCache *cache.Cache, pokeClient *pokeapi.Client, parameter string, pokedex map[string]pokeapi.Pokemon) error {
	url := ""
	if configuration.Previous != "" {
		url = configuration.Previous
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
