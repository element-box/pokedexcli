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
				err := command.callback(cfg)
				if err != nil {
					fmt.Println(err)
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
	nextLocURL    *string
	prevLocURL    *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a message",
			callback:    displayHelp,
		},
		"map": {
			name:        "map",
			description: "Displays 20 location areas of Pokemon",
			callback:    displayMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displsy previous 20 location areas",
			callback:    displayMapB,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
