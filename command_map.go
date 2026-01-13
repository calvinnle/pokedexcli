package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaUrl)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}

	listLocations := resp.Results

	for i := range listLocations {
		fmt.Printf("%v\n", listLocations[i].Name)
	}

	cfg.NextPage()
	cfg.PrintCurrentPage()

	cfg.nextLocationAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationAreaUrl == nil {
		return fmt.Errorf("you're on the first page dumbass")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaUrl)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}

	listLocations := resp.Results

	for i := range listLocations {
		fmt.Printf("%v\n", listLocations[i].Name)
	}

	cfg.PrevPage()
	cfg.PrintCurrentPage()

	cfg.nextLocationAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous
	return nil
}
