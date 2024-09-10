package pokeapi

import (
	"math/rand"
)

func (c *Client) TryCatchPokemon(name string) (string, error) {
	pokemon, err := c.GetPokemon(name)
	if err != nil {
		return "", err
	}

	if weightedRandomCatch(pokemon.BaseExperience) {
		c.pokedex.Add(pokemon)
		return pokemon.Name + " was caught!", nil
	}

	return pokemon.Name + " escaped!", nil
}

func weightedRandomCatch(weight int) bool {
	normalizedWeight := (float64(weight) - 30.0) / (220.0 - 30.0)
	randomVal := rand.Float64()
	threshold := 0.5 + (normalizedWeight * 0.5)

	return randomVal >= threshold
}
