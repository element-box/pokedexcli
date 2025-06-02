package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (c *Client) ListExploreLocation(location *string) (PokemonEncountersResponse, error) {
	url := baseURL + "/location-area"
	locArea, err := c.ListLocations(&url)
	if err != nil {
		return PokemonEncountersResponse{}, err
	}

	locationURL := ""
	for _, loc := range *&locArea.Results {
		if strings.Compare(loc.Name, *location) == 0 {
			locationURL = loc.URL
			break
		}
	}

	if locationURL == "" {
		fmt.Println("Locaiton not found!")
		return PokemonEncountersResponse{}, nil
	}
	val, exists := c.cache.Get(url)
	pokemonEncounters := PokemonEncountersResponse{}
	if !exists {
		if locationURL != "" {
			res, err := http.Get(locationURL)
			if err != nil {
				return PokemonEncountersResponse{}, err
			}
			body, err := io.ReadAll(res.Body)
			if err != nil {
				return PokemonEncountersResponse{}, err
			}
			defer res.Body.Close()

			c.cache.Add(locationURL, body)

			err = json.Unmarshal(body, &pokemonEncounters)
			if err != nil {
				return PokemonEncountersResponse{}, err
			}
		}
	} else {
		err := json.Unmarshal(val, &pokemonEncounters)
		if err != nil {
			return PokemonEncountersResponse{}, err
		}
	}
	return pokemonEncounters, nil
}
