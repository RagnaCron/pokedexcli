package pokecommand

import (
	"github.com/RagnaCron/pokedexcli/internal/pokeapi"
	"github.com/RagnaCron/pokedexcli/internal/pokecache"
)

type Config struct {
	Previous *string
	Next     *string
	Cache    *pokecache.Cache
	Pokedex  map[string]pokeapi.Pokemon
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
		"catch": {
			Name:        "catch <pokemon_name>",
			Description: "Catch a Pokemon",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect <pokemon_name>",
			Description: "Inspect a caught Pokemon in the Pokedex",
			Callback:    commandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "See all the Pokemon you've caught",
			Callback:    commandPokedex,
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
