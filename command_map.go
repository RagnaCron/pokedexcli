package main

import (
	"errors"

	"github.com/RagnaCron/pokedexcli/internal/pokeapi"
)

func commandMap(conf *cliConfig) error {
	resource, err := pokeapi.GetLocations(conf.next, conf.cache)
	if err != nil {
		return err
	}

	conf.next = resource.Next
	conf.previous = resource.Previous

	resource.PrintResults()

	return nil
}

func commandMapb(conf *cliConfig) error {
	if conf.previous == nil {
		return errors.New("you're on the first page")
	}
	resource, err := pokeapi.GetLocations(conf.previous, conf.cache)
	if err != nil {
		return err
	}

	conf.next = resource.Next
	conf.previous = resource.Previous

	resource.PrintResults()

	return nil
}
