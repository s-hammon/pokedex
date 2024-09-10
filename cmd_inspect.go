package main

import "fmt"

func cmdInspect(cfg *pokeConfig, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: inspect <pokemon>")
	}

	pokemonName := args[0]
	if pokemon, ok := cfg.pokeClient.Pokedex.Get(pokemonName); ok {
		fmt.Printf("%v", pokemon)
		return nil
	}

	fmt.Println("you have not caught that pokemon yet")
	return nil
}
