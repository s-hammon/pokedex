package main

import (
	"fmt"
)

func cmdMap(cfg *pokeConfig) error {
	areas, err := cfg.pokeClient.GetLocations(cfg.nextURL)
	if err != nil {
		return err
	}

	cfg.nextURL = areas.Next
	cfg.prevURL = areas.Previous

	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func cmdMapB(cfg *pokeConfig) error {
	if cfg.prevURL == nil {
		return fmt.Errorf("no previous page")
	}

	areas, err := cfg.pokeClient.GetLocations(cfg.prevURL)
	if err != nil {
		return err
	}

	cfg.nextURL = areas.Next
	cfg.prevURL = areas.Previous

	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}

	return nil
}
