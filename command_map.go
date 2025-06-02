package main

import (
	"fmt"
)

func commandMapF(cfg *config, _ string) error {
	locAreaRes, err := cfg.pokeapiClient.ListLocations(cfg.nextLocURL)
	if err != nil {
		return err
	}
	cfg.nextLocURL = locAreaRes.Next
	cfg.prevLocURL = locAreaRes.Previous

	for _, loc := range locAreaRes.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapB(cfg *config, _ string) error {
	if cfg.prevLocURL == nil {
		fmt.Println("At the start of the list! Try 'map' command")
		return nil
	}

	locAreaRes, err := cfg.pokeapiClient.ListLocations(cfg.prevLocURL)
	if err != nil {
		return err
	}

	cfg.nextLocURL = locAreaRes.Next
	cfg.prevLocURL = locAreaRes.Previous

	for _, loc := range locAreaRes.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
