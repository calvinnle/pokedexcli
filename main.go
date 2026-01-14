package main

import (
	"fmt"
	"time"

	"github.com/calvinnle/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       *pokeapi.Client
	nextLocationAreaUrl *string
	prevLocationAreaUrl *string
	currentPage         int
	pokedex             []string
}

func (c *config) NextPage() {
	c.currentPage++
}

func (c *config) PrevPage() {
	c.currentPage--
}

func (c *config) PrintCurrentPage() {
	fmt.Printf("You're on page %d\n", c.currentPage)
}

func main() {
	currentpage := 0
	interval := time.Hour
	cfg := config{
		pokeapiClient: pokeapi.NewClient(interval),
		currentPage:   currentpage,
        pokedex: make([]string, 0),
	}

	startRepl(&cfg)
}
