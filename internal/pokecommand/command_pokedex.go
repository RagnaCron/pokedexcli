package pokecommand

import (
	"errors"
	"fmt"
)

func commandPokedex(conf *Config, args ...string) error {
	if len(conf.Pokedex) == 0 {
		return errors.New("you have to catch some some Pokemon")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range conf.Pokedex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
