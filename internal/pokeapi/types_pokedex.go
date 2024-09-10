package pokeapi

type Pokedex struct {
	CaughtPokemon map[string]Pokemon
}

func (p *Pokedex) Add(pokemon Pokemon) {
	p.CaughtPokemon[pokemon.Name] = pokemon
}
