package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rodriguesfrancisco/go-pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(pageConfig *pageConfig) error
}

type pageConfig struct {
	Next       string
	Previous   string
	Cache      pokecache.Cache
	Parameters []string
	Pokedex    map[string]PokemonResponse
}

type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var pageConfigInstance = &pageConfig{
	Next:     "",
	Previous: "",
	Pokedex:  make(map[string]PokemonResponse),
	Cache:    *pokecache.NewCache(2 * time.Minute),
}

var availableCommands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the pokedex",
		callback:    commandExit,
	},
	"map": {
		name:        "map",
		description: "Returns the next 20 location areas in Pokemon world",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Returns the previous 20 location areas in Pokemon world",
		callback:    commandMapb,
	},
	"explore": {
		name:        "explore",
		description: "Receives an location area and prints all pokemons on that area",
		callback:    commandExplore,
	},
	"catch": {
		name:        "catch",
		description: "Receives an pokémon name and tries to catch it",
		callback:    commandCatch,
	},
	"inspect": {
		name:        "inspect",
		description: "Returns informations about a pokémon in your pokedex",
		callback:    commandInspect,
	},
	"pokedex": {
		name:        "pokedex",
		description: "Lists all pokémons in your pokedex",
		callback:    commandPokedex,
	},
}

func cleanInput(input string) []string {
	words := strings.Fields(input)
	cleanedWords := []string{}
	for _, word := range words {
		trimmed := strings.TrimSpace(word)
		lowered := strings.ToLower(trimmed)
		cleanedWords = append(cleanedWords, lowered)
	}
	return cleanedWords
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		commands := cleanInput(input)

		if commands[0] == "help" {
			commandHelp(pageConfigInstance)
			continue
		}

		currentCommand, ok := availableCommands[commands[0]]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		pageConfigInstance.Parameters = commands[1:]

		err := currentCommand.callback(pageConfigInstance)
		if err != nil {
			fmt.Println(err)
		}
	}
}
