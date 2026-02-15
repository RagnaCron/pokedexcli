package pokecommand

import (
	"errors"

	"github.com/RagnaCron/pokedexcli/internal/pokeapi"
)

func commandMapf(conf *Config, args ...string) error {
	resource, err := pokeapi.GetLocations(conf.Next, conf.Cache)
	if err != nil {
		return err
	}

	conf.Next = resource.Next
	conf.Previous = resource.Previous

	resource.PrintResults()

	return nil
}

func commandMapb(conf *Config, args ...string) error {
	if conf.Previous == nil {
		return errors.New("you're on the first page")
	}
	resource, err := pokeapi.GetLocations(conf.Previous, conf.Cache)
	if err != nil {
		return err
	}

	conf.Next = resource.Next
	conf.Previous = resource.Previous

	resource.PrintResults()

	return nil
}
