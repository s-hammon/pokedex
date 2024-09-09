package main

import (
	"time"

	"github.com/s-hammon/pokedex/internal/pokeapi"
)

var proompt string = "pokedex"

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	cfg := &pokeConfig{
		pokeClient: &client,
	}
	startRepl(cfg)
}
