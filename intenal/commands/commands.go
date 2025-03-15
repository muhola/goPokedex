package commands

type cliCommand struct {
	name        string
	description string
	callback    func() error
	config      config
}

type config struct {
	Next     string
	Previous string
}

var cliComands map[string]cliCommand

func getCommands() map[string]cliCommand {
	return cliComands
}

func initComands() {
	cliComands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: " Display a help message",
			callback:    commandHelp,
			config:      config{},
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world ",
			callback:    commandMap,
		},
		"exit": {
			name:        "exit",
			description: " Exit the poked",
			callback:    commandExit,
		},
	}
}
