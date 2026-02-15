package main

import (
	"errors"
	"fmt"

	"github.com/RagnaCron/pokedexcli/internal/pokeapi"
)

func commandExplore(conf *cliConfig, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location name")
	}
	name := args[0]
	fmt.Printf("Exploring %v...\n", name)

	err := pokeapi.ShowEncounters(conf.cache, name)
	if err != nil {
		return err
	}

	return nil
}
