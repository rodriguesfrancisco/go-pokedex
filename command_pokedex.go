package main

import "fmt"

func commandPokedex(page *pageConfig) error {
	fmt.Println("Your pokedex:")
	for _, pokemon := range page.Pokedex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
