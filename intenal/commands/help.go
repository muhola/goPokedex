package commands

import "fmt"

func commandHelp() error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage: \n\n")

	for _, value := range commands {
		fmt.Println(value.name, ": ", value.description)
	}
	fmt.Printf("\n")
	return nil
}
