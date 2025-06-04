package main

import (
	"fmt"
)

func commandInspect(cfg *config, pokemon string) error {
	pokemonInfo, exists := cfg.pokedex[pokemon]
	if exists {
		fmt.Printf("Name: %s\n", pokemonInfo.Name)
		fmt.Printf("Height: %d\n", pokemonInfo.Height)
		fmt.Printf("Weight: %d\n", pokemonInfo.Weight)
		fmt.Println("Stats:")
		for _, stats := range pokemonInfo.Stats {
			fmt.Printf("  -%s: %d\n", stats.Stat.Name, stats.BaseStat)
		}
		fmt.Println("Types:")
		for _, pokeType := range pokemonInfo.Types {
			fmt.Printf("  - %s\n", pokeType.Type.Name)
		}
	} else {
		fmt.Printf("You haven't caught %s yet!", pokemon)
	}
	return nil
}
