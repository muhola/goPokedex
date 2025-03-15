package commands

import (
	"fmt"

	"github.com/muhola/goPokedex/internal/pokeapi"
)

func commandMap(configuration *configuration) error {
	url := ""
	if configuration.Next != "" {
		url = configuration.Next
	}
	locations := pokeapi.GetLocation(url)
	if locations.Next != "<nil>" {
		configuration.Next = locations.Next

	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}
