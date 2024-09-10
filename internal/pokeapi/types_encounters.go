package pokeapi

type Encounters struct {
	AreaID            int                `json:"id"`
	AreaName          string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}
