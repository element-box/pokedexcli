package main

import (
	"pokedexcli/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokedexMap := make(map[string]pokeapi.Pokemon)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       pokedexMap,
	}
	// REPL - Read Eval Print Loop
	startRepl(cfg)
}
