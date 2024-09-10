package main

import "fmt"

func cmdExplore(cfg *pokeConfig, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: explore <area>")
	}

	areaName := args[0]
	encounters, err := cfg.pokeClient.GetEncounters(areaName)
	if err != nil {
		return err
	}

	for _, encounter := range encounters.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
