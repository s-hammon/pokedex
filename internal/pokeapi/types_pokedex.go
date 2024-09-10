package pokeapi

type Pokedex struct {
	CaughtPokemon map[string]Pokemon
}

func (p *Pokedex) Add(pokemon Pokemon) {
	p.CaughtPokemon[pokemon.Name] = pokemon
}

func (p *Pokedex) Get(name string) (Pokemon, bool) {
	pokemon, ok := p.CaughtPokemon[name]
	return pokemon, ok
}

func (p *Pokedex) GetAll() []Pokemon {
	pokemon := make([]Pokemon, 0, len(p.CaughtPokemon))
	for _, p := range p.CaughtPokemon {
		pokemon = append(pokemon, p)
	}
	return pokemon
}
