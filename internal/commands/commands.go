package commands

import (
	"fmt"
	"strings"

	"github.com/muhola/goPokedex/internal/cache"
	"github.com/muhola/goPokedex/internal/pokeapi"
)

var sharedConfig = &Configuration{}
var cliComands map[string]CliCommand

func GetCommands() map[string]CliCommand {
	return cliComands
}

func GetCommand(command, parameter string, cache *cache.Cache, pokeClient *pokeapi.Client, pokedex map[string]pokeapi.Pokemon) {
	clicmd, ok := cliComands[strings.ToLower(command)]
	if ok {
		cmdError := clicmd.callback(clicmd.config, cache, pokeClient, parameter, pokedex)
		if cmdError != nil {
			fmt.Println(cmdError.Error())
		}

	} else {
		fmt.Println("Unknown command")
	}

}

func InitComands() {

	cliComands = map[string]CliCommand{
		"help": {
			name:        "help",
			description: "Display this help message",
			callback:    commandHelp,
			config:      sharedConfig,
		},
		"catch": {
			name:        "catch",
			description: "Try and catch aa pokemon, Catch <Pokemon id/name>",
			callback:    commandCatch,
			config:      sharedConfig,
		},
		"explore": {
			name:        "explore",
			description: "Display Pokemons in the location Area, Explore <area id/name>",
			callback:    commandExplore,
			config:      sharedConfig,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect your pokedex, Inspect <Pokemon name>",
			callback:    commandIncpect,
			config:      sharedConfig,
		},
		"map": {
			name:        "map",
			description: "Displays Next 20 location areas in the Pokemon world",
			callback:    commandMap,
			config:      sharedConfig,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the Previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
			config:      sharedConfig,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all captrued pokemons",
			callback:    commandPokedex,
			config:      sharedConfig,
		},
		"exit": {
			name:        "exit",
			description: "Exit the poked",
			callback:    commandExit,
			config:      sharedConfig,
		},
	}
}
