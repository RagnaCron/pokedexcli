package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Result struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Resource struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Result `json:"results"`
}

func GetMap(url string) (Resource, error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	res, err := http.Get(url)
	if err != nil {
		return Resource{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return Resource{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return Resource{}, err
	}

	resource := Resource{}
	err = json.Unmarshal(body, &resource)
	if err != nil {
		return Resource{}, err
	}

	return resource, nil
}
