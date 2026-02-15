package main

import "github.com/RagnaCron/pokedexcli/internal/pokecache"

type cliConfig struct {
	previous *string
	next     *string
	cache    *pokecache.Cache
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
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
	}
}
