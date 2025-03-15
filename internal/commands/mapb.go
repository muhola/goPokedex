package commands

import (
	"fmt"

	"github.com/muhola/goPokedex/internal/pokeapi"
)

func commandMapb(configuration *configuration) error {
	url := ""
	if configuration.Previous != "" {
		url = configuration.Previous
	}
	locations := pokeapi.GetLocation(url)
	if locations.Previous != "<nil>" {
		configuration.Previous = locations.Previous
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}
