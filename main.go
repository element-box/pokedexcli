package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func displayHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	// for _, command := range commands {
	// 	fmt.Printf("%v: %v", command.name, command.description)
	// }
	return nil
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Displays a message",
		callback:    displayHelp,
	},
}

func cleanInput(text string) []string {
	slice := strings.Fields(strings.ToLower(text))
	return slice
}

func main() {
	// REPL - Read Eval Print Loop
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := scanner.Text()
			sanInput := cleanInput(userInput)
			if len(sanInput) == 0 {
				continue
			}
			command, exists := commands[sanInput[0]]
			if exists {
				command.callback()
			} else {
				fmt.Println("Unknown command")
			}
		} else {
			return
		}
	}
}
