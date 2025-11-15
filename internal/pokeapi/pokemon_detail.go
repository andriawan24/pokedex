package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) PokemonDetail(name string) (Pokemon, error) {
	fullUrl := baseURL + "/pokemon/" + name

	pokemonResponse := Pokemon{}

	if v, ok := c.cache.Get(fullUrl); ok {
		err := json.Unmarshal(v, &pokemonResponse)
		if err != nil {
			return pokemonResponse, err
		}

		return pokemonResponse, nil
	}

	res, err := http.Get(fullUrl)
	if err != nil {
		return pokemonResponse, err
	}

	if res.StatusCode == http.StatusNotFound {
		return pokemonResponse, fmt.Errorf("%s is not found", name)
	}

	if res.StatusCode != http.StatusOK {
		return pokemonResponse, fmt.Errorf("request is not success %v", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return pokemonResponse, err
	}

	if err = json.Unmarshal(data, &pokemonResponse); err != nil {
		return pokemonResponse, err
	}

	c.cache.Add(fullUrl, data)

	return pokemonResponse, nil
}
