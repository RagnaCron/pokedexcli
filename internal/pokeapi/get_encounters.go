package pokeapi

import (
	"encoding/json"
	"fmt"

	"github.com/RagnaCron/pokedexcli/internal/pokecache"
)

type Encounters struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (e *Encounters) ShowPokemons() {
	fmt.Println("Found Pokemon:")
	for _, encounter := range e.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
}

func ShowEncounters(cache *pokecache.Cache, name string) error {
	url := baseURL + "/location-area/" + name + "/"

	encounters, err := getEncounters(cache, url)
	if err != nil {
		return err
	}

	encounters.ShowPokemons()

	return nil
}

func getEncounters(cache *pokecache.Cache, url string) (Encounters, error) {
	body, err := Get(url, cache)
	if err != nil {
		return Encounters{}, err
	}

	var encoutners Encounters
	err = json.Unmarshal(body, &encoutners)
	if err != nil {
		return Encounters{}, err
	}

	return encoutners, nil
}
