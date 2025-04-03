package commands

import (
	"fmt"
	"strings"

	"github.com/muhola/goPokedex/internal/cache"
	"github.com/muhola/goPokedex/internal/pokeapi"
)

var sharedConfig = &Configuration{}
var cliComands map[string]CliCommand
var pokeCache *cache.Cache

func GetCommands() map[string]CliCommand {
	return cliComands
}

func GetCommand(input string, cache *cache.Cache, pokeClient *pokeapi.Client) {
	clicmd, ok := cliComands[strings.ToLower(input)]
	if ok {
		cmdError := clicmd.callback(clicmd.config, cache, pokeClient)
		if cmdError != nil {
			fmt.Errorf(cmdError.Error())
		}

	} else {
		fmt.Println("Unknown command")
	}

}

func InitComands() {

	cliComands = map[string]CliCommand{
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
