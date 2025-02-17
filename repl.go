package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/JLee871/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
	Pokedex       map[string]pokeapi.Pokemon
}

// Commands mapping
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Gets the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Gets the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location name>",
			description: "Returns pokemon found at location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon name>",
			description: "Attempts to catch pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon name",
			description: "shows data on pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "shows your caught pokemon",
			callback:    commandPokedex,
		},
	}
}

// Read Eval Print Loop
func startRepl(c *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		command := words[0]
		parameter := ""

		if len(words) > 1 {
			parameter = words[1]
		}

		comm, ok := getCommands()[command]
		if ok {
			err := comm.callback(c, parameter)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

// Splits input on whitespace and makes lowercase
func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
