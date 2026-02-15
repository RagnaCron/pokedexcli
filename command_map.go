package main

import (
	"errors"

	"github.com/RagnaCron/pokedexcli/internal/pokeapi"
)

func commandMapf(conf *cliConfig, args ...string) error {
	resource, err := pokeapi.GetLocations(conf.next, conf.cache)
	if err != nil {
		return err
	}

	conf.next = resource.Next
	conf.previous = resource.Previous

	resource.PrintResults()

	return nil
}

func commandMapb(conf *cliConfig, args ...string) error {
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
