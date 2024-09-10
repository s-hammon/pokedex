package main

import "fmt"

func cmdCatch(cfg *pokeConfig, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: catch <pokemon>")
	}

	pokemonName := args[0]
	fmt.Printf("Trying to catch %s...\n", pokemonName)

	result, err := cfg.pokeClient.TryCatchPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}
