package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := scanner.Text()
			sanInput := cleanInput(userInput)
			if len(sanInput) == 0 {
				continue
			}
			command, exists := getCommands()[sanInput[0]]
			if exists {
				if len(sanInput) == 2 {
					err := command.callback(cfg, sanInput[1])
					if err != nil {
						fmt.Println(err)
					}
				} else {
					err := command.callback(cfg, "")
					if err != nil {
						fmt.Println(err)
					}
				}
				continue
			} else {
				fmt.Println("Unknown command")
				continue
			}
		} else {
			return
		}
	}
}

func cleanInput(text string) []string {
	slice := strings.Fields(strings.ToLower(text))
	return slice
}

type config struct {
	pokeapiClient pokeapi.Client
	pokedex       map[string]pokeapi.Pokemon
	nextLocURL    *string
	prevLocURL    *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a message",
			callback:    commandHelp,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokeball",
			callback:    commandCatch,
		},
		"explore": {
			name:        "explore",
			description: "Display possible Pokemon encounters for a given area",
			callback:    commandExplore,
		},
		"inspect": {
			name:        "inspect",
			description: "Dislay the data for a pokemon that's been caught",
			callback:    commandInspect,
		},
		"map": {
			name:        "map",
			description: "Displays 20 location areas of Pokemon",
			callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 location areas",
			callback:    commandMapB,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display your caught Pokemon",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
