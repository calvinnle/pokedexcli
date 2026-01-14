package main

import (
    "fmt"
)

func commandPokedex(cfg *config, _ string) error {
    if len(cfg.pokedex) == 0 {
        fmt.Println("You has caught none! Stop being a lazy ass and go get some!")
    }
    for _, pokemon := range cfg.pokedex {
        fmt.Println("You has caught: ")
        fmt.Printf("- %s\n", pokemon)
    }

	return nil
}
