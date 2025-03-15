package commands

import (
	"fmt"
	"os"
)

func commandExit(config *configuration) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return nil
}
