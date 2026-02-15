package pokeapi

import (
	"encoding/json"
	"fmt"

	"github.com/RagnaCron/pokedexcli/internal/pokecache"
)

type Locations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (r Locations) PrintResults() {
	for _, location := range r.Results {
		fmt.Println(location.Name)
	}
}

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
