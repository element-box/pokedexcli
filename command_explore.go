package main

import (
	"fmt"
)

func commandExplore(cfg *config, location string) error {
	if location == "" {
		fmt.Println("The `explore` command needs a location! Try `explore mt-coronet-2f`")
		return nil
	}
	pokemonEncounters, err := cfg.pokeapiClient.ListExploreLocation(&location)
	if err != nil {
		return err
	}

	for _, pokemon := range pokemonEncounters.PokemonEncounters {
		fmt.Printf("- %v\n", pokemon.Pokemon.Name)
	}

	return nil
}
