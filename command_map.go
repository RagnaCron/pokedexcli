package main

import (
	"fmt"

	"github.com/RagnaCron/pokedexcli/internal/pokeapi"
)

func commandMap(conf *cliConfig) error {
	resource, err := pokeapi.GetMap(conf.next)
	if err != nil {
		return err
	}

	if resource.Next != nil {
		conf.next = *resource.Next
	}
	if resource.Previous != nil {
		conf.previous = *resource.Previous
	}

	for _, location := range resource.Results {
		fmt.Println(location.Name)
	}

	return nil
}
