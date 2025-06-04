package main

import (
	"fmt"
	"math"
	"math/rand"
)

func commandCatch(cfg *config, pokemon string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	pokemonData, err := cfg.pokeapiClient.GetPokemon(pokemon)
	if err != nil {
		return err
	}

	baseChance := .8
	baseExp := pokemonData.BaseExperience
	totalChance := (baseChance * 100) / float64(baseExp)
	randomChance := float64(rand.Intn(100)) / 100

	totalChance = math.Min(1.0, totalChance)

	if randomChance <= totalChance {
		fmt.Printf("%s was caught!\n", pokemon)
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	cfg.pokedex[pokemon] = pokemonData

	return nil
}
