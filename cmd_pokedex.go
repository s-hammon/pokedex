package main

import "fmt"

func cmdPokedex(cfg *pokeConfig, args ...string) error {
	if len(args) != 0 {
		return fmt.Errorf("usage: pokedex")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokeClient.Pokedex.GetAll() {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
