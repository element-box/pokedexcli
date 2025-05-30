package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreaResponse struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func _displayLocation(url string, cfg *config) error {

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return err
	}
	locAreaRes := LocationAreaResponse{}
	err = json.Unmarshal(body, &locAreaRes)
	if err != nil {
		return err
	}
	cfg.Next = locAreaRes.Next
	cfg.Previous = locAreaRes.Previous
	for _, loc := range locAreaRes.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func displayMap(cfg *config) error {
	url := ""
	if cfg.Next != nil && *cfg.Next != "" {
		url = *cfg.Next
	} else {
		url = "https://pokeapi.co/api/v2/location-area/"
	}
	return _displayLocation(url, cfg)
}

func displayMapB(cfg *config) error {
	url := ""
	if cfg.Previous != nil && *cfg.Previous != "" {
		url = *cfg.Previous
	} else {
		fmt.Println("At the beginning of the locations list! Try using 'map' command")
		return nil
	}

	return _displayLocation(url, cfg)
}
