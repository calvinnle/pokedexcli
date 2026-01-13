package pokeapi

import (
	"net/http"
	"time"

	"github.com/calvinnle/pokedexcli/internal/pokecache"
)

const BaseURL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient  *http.Client
	clientCache pokecache.Cache
}

func NewClient(interval time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Minute,
		},
        clientCache: *pokecache.NewCache(interval),
	}
}
