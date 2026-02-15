package pokeapi

import (
	"encoding/json"
)

func GetMap(url string) (Locations, error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	body, err := Get(url)
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
