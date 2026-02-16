package pokecommand

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/RagnaCron/pokedexcli/internal/pokeapi"
)

func commandCatch(conf *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}
	name := args[0]
	fmt.Printf("Throwing a Pokeball at %v...\n", name)

	pokemon, err := pokeapi.GetPokemon(conf.Cache, name)
	if err != nil {
		return err
	}

	catch(conf, &pokemon)

	return nil
}

func catch(conf *Config, pokemon *pokeapi.Pokemon) {
	canCatch := attempt(pokemon.BaseExperience)
	if canCatch {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Println("You may now inspect it with the inspect command.")
		conf.Pokedex[pokemon.Name] = *pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

}

func attempt(baseExp int) bool {
	const k = 0.01 // can be tuned

	p := 1.0 / (1.0 + k*float64(baseExp))

	if p < 0.05 {
		p = 0.05
	}

	return newR().Float64() < p
}

func newR() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
