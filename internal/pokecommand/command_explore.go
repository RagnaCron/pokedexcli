package pokecommand

import (
	"errors"
	"fmt"

	"github.com/RagnaCron/pokedexcli/internal/pokeapi"
)

func commandExplore(conf *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}
	name := args[0]
	fmt.Printf("Exploring %v...\n", name)

	err := pokeapi.ShowEncounters(conf.Cache, name)
	if err != nil {
		return err
	}

	return nil
}
