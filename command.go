package main

type cliConfig struct {
	previous string
	next     string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*cliConfig) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
	}
}
