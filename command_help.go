package main

import (
	"fmt"
)

func commandHelp(_ *pageConfig) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Printf("%s: %s\n", "help", "Displays a help message")
	for _, command := range availableCommands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
