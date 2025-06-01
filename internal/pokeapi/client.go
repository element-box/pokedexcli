package pokeapi

import (
	"net/http"
	"pokedexcli/internal/pokecache"
	"time"
)

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(10 * time.Second),
	}
}
