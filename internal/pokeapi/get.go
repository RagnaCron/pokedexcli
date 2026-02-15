package pokeapi

import (
	"fmt"
	"io"
	"net/http"

	"github.com/RagnaCron/pokedexcli/internal/pokecache"
)

func Get(url string, cache *pokecache.Cache) ([]byte, error) {
	body, ok := cache.Get(url)
	if ok {
		return body, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return nil, err
	}

	cache.Add(url, body)

	return body, nil
}
