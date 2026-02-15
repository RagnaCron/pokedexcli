package pokecommand

import "github.com/RagnaCron/pokedexcli/internal/pokecache"

type Config struct {
	Previous *string
	Next     *string
	Cache    *pokecache.Cache
}

type Command struct {
	Name        string
	Description string
	Callback    func(*Config, ...string) error
}

func Get() map[string]Command {
	return map[string]Command{
		"help": {
			Name:        "help",
			Description: "Display a help message",
			Callback:    commandHelp,
		},
		"explore": {
			Name:        "explore <location_name>",
			Description: "Explore a location",
			Callback:    commandExplore,
		},
		"map": {
			Name:        "map",
			Description: "Get the next page of locations",
			Callback:    commandMapf,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Get the previous page of locations",
			Callback:    commandMapb,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
	}
}
