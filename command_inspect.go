package main

import "fmt"

func commandInspect(cfg *config, arg string) error {
	minimalPokemon, err := cfg.pokeapiClient.GetPokemonMinimal(arg)
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}

	fmt.Printf("Name: %s (#%d)\n", minimalPokemon.Name, minimalPokemon.ID)
	fmt.Printf("Height: %.1f m\n", minimalPokemon.Height)
	fmt.Printf("Weight: %.1f kg\n", minimalPokemon.Weight)
	fmt.Printf("Base Experience: %d\n", minimalPokemon.BaseExperience)

	// Types
	fmt.Print("Types: ")
	for i, t := range minimalPokemon.Types {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(t)
	}
	fmt.Println()

	// Stats
	fmt.Println("Stats:")
	for name, value := range minimalPokemon.Stats {
		fmt.Printf("  %s: %d\n", name, value)
	}

	// Abilities
	fmt.Println("Abilities:")
	for _, a := range minimalPokemon.Abilities {
		fmt.Printf("  • %s", a.Name)
		if a.IsHidden {
			fmt.Print(" (hidden)")
		}
		fmt.Println()
	}

	return nil
}
