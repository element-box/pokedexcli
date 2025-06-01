package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	val, exists := c.cache.Get(url)
	locAreaRes := LocationAreaResponse{}
	if !exists {
		res, err := http.Get(url)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		defer res.Body.Close()

		c.cache.Add(url, body)

		err = json.Unmarshal(body, &locAreaRes)
		if err != nil {
			return LocationAreaResponse{}, err
		}
	} else {
		err := json.Unmarshal(val, &locAreaRes)
		if err != nil {
			return LocationAreaResponse{}, err
		}
	}
	return locAreaRes, nil
}
