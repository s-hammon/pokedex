package pokeapi

import (
	"net/http"
	"time"

	"github.com/s-hammon/pokedex/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient *http.Client
	Pokedex    *Pokedex
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: &http.Client{
			Timeout: timeout,
		},
		Pokedex: &Pokedex{
			CaughtPokemon: make(map[string]Pokemon),
		},
	}
}
