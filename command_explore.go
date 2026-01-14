package main

import "fmt"


func commandExplore(cfg *config, location string) error {
    if len(location) == 0 {
        return fmt.Errorf("location not provided - empty")
    }
    resp, err := cfg.pokeapiClient.GetLocationArea(location)
    if err != nil {
        return err
    }

    pokemons := resp.PokemonEncounters

    fmt.Printf("Pokemon in %s:\n", location)
    for _, pokemon := range pokemons {
        fmt.Printf("%v\n", pokemon.Pokemon.Name)
    }

    return nil
}
