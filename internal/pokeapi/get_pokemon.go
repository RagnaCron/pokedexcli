package pokeapi

import (
	"encoding/json"

	"github.com/RagnaCron/pokedexcli/internal/pokecache"
)

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

func GetPokemon(cache *pokecache.Cache, name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name + "/"

	body, err := Get(url, cache)
	if err != nil {
		return Pokemon{}, err
	}

	var pokemon Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
