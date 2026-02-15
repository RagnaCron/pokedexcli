package pokeapi

import (
	"encoding/json"

	"github.com/RagnaCron/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

func GetLocations(pageURL *string, cache *pokecache.Cache) (Locations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	body, err := Get(url, cache)
	if err != nil {
		return Locations{}, err
	}

	resource := Locations{}
	err = json.Unmarshal(body, &resource)
	if err != nil {
		return Locations{}, err
	}

	return resource, nil
}
