package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemon string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon

	val, exists := c.cache.Get(url)
	pokemonExp := Pokemon{}
	if !exists {
		res, err := http.Get(url)
		if err != nil {
			return Pokemon{}, err
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()

		c.cache.Add(url, body)

		err = json.Unmarshal(body, &pokemonExp)
		if err != nil {
			return Pokemon{}, err
		}
	} else {
		err := json.Unmarshal(val, &pokemonExp)
		if err != nil {
			return Pokemon{}, err
		}
	}

	return pokemonExp, nil
}
