package commands

import (
	"fmt"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*configuration) error
	config      *configuration
}

type configuration struct {
	Next     string
	Previous string
}

var sharedConfig = &configuration{}
var cliComands map[string]cliCommand

func GetCommands() map[string]cliCommand {
	return cliComands
}

func GetCommand(input string) {
	clicmd, ok := cliComands[strings.ToLower(input)]
	if ok {
		cmdError := clicmd.callback(clicmd.config)
		if cmdError != nil {
			fmt.Errorf(cmdError.Error())
		}

	} else {
		fmt.Println("Unknown command")
	}

}

func InitComands() {

	cliComands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: " Display a help message",
			callback:    commandHelp,
			config:      sharedConfig,
		},
		"map": {
			name:        "map",
			description: "Displays Next 20 location areas in the Pokemon world ",
			callback:    commandMap,
			config:      sharedConfig,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the Previous 20 location areas in the Pokemon world ",
			callback:    commandMapb,
			config:      sharedConfig,
		},
		"exit": {
			name:        "exit",
			description: " Exit the poked",
			callback:    commandExit,
			config:      sharedConfig,
		},
	}
}
