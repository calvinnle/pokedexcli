package main

import "github.com/calvinnle/pokedexcli/internal/pokeapi"

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaUrl *string
	prevLocationAreaUrl *string
}

func main() {
    cfg := config{
        pokeapiClient: *pokeapi.NewClient(),
    }

	startRepl(&cfg)
}
