package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationDetail(LocationArea string) (LocationAreaDetail, error) {
	fullUrl := baseURL + "/location-area/" + LocationArea

	if v, ok := c.cache.Get(fullUrl); ok {
		locationResponse := LocationAreaDetail{}

		err := json.Unmarshal(v, &locationResponse)
		if err != nil {
			return locationResponse, err
		}

		return locationResponse, nil
	}

	res, err := http.Get(fullUrl)
	if err != nil {
		return LocationAreaDetail{}, err
	}

	if res.StatusCode != http.StatusOK {
		return LocationAreaDetail{}, fmt.Errorf("request is not success %v", res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaDetail{}, err
	}

	var locationResponse LocationAreaDetail
	if err := json.Unmarshal(data, &locationResponse); err != nil {
		return LocationAreaDetail{}, err
	}

	c.cache.Add(fullUrl, data)

	return locationResponse, nil
}
