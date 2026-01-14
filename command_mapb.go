package main

import "fmt"

func commandMapb(cfg *config, arg string) error {
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
