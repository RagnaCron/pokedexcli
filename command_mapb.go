package main

import "github.com/RagnaCron/pokedexcli/internal/pokeapi"

func commandMapb(conf *cliConfig) error {
	resource, err := pokeapi.GetMap(conf.previous)
	if err != nil {
		return err
	}

	if resource.Next != nil {
		conf.next = *resource.Next
	}
	if resource.Previous != nil {
		conf.previous = *resource.Previous
	}

	resource.PrintResults()

	return nil
}
