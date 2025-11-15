package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocation(pageURL *string) (LocationArea, error) {
	fullUrl := baseURL + "/location-area?offset=0&limit=20"
	if pageURL != nil {
		fullUrl = *pageURL
	}

	if v, ok := c.cache.Get(fullUrl); ok {
		locationResponse := LocationArea{}

		err := json.Unmarshal(v, &locationResponse)
		if err != nil {
			return locationResponse, err
		}

		return locationResponse, nil
	}

	res, err := http.Get(fullUrl)
	if err != nil {
		return LocationArea{}, err
	}

	if res.StatusCode != http.StatusOK {
		return LocationArea{}, fmt.Errorf("request is not success %v", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	var locationResponse LocationArea
	if err := json.Unmarshal(data, &locationResponse); err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullUrl, data)

	return locationResponse, nil
}
