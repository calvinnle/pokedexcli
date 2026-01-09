package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas()
	if err != nil {
        return fmt.Errorf("error making request: %v", err)
	}

    listLocations := resp.Results

    for i := range listLocations {
        fmt.Printf("Location: %v\n", listLocations[i].Name)
    }

    return nil
}
