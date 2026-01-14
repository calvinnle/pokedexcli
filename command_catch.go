package main

import (
	"fmt"
)

func commandCatch(cfg *config, arg string) error {
	caught, err := cfg.pokeapiClient.CatchPokemon(arg)
    if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}

    if caught {
        cfg.pokedex = append(cfg.pokedex, arg)
    }

	return nil
}
