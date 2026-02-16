package pokecommand

import (
	"errors"
)

func commandInspect(conf *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]

	pokemon, ok := conf.Pokedex[name]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	pokemon.ShowStats()

	return nil
}
