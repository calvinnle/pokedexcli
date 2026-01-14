package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonDetails(name string) (Pokemon, error) {
	var pokemon Pokemon
	if len(name) == 0 {
		return pokemon, fmt.Errorf("name is not defined")
	}

	// store full url to make use of caching
	endpoint := "/pokemon/" + name
	fullUrl := BaseURL + endpoint

	// cache hit
	data, ok := c.clientCache.Get(fullUrl)
	if ok {
		fmt.Println("cache hit!")
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return pokemon, err
		}

		return pokemon, nil
	}

	fmt.Println("cache miss!")
	// cache miss
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return pokemon, fmt.Errorf("error creating new request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return pokemon, fmt.Errorf("error making request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return pokemon, fmt.Errorf("bad status: %d %s",
			resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return pokemon, fmt.Errorf("Error reading body")
	}

	if err := json.Unmarshal(data, &pokemon); err != nil {
		return pokemon, err
	}

	// cache the data
	c.clientCache.Set(fullUrl, data)

	return pokemon, nil
}

func (p *Pokemon) ToMinimal() MinimalPokemon {
	minimal := MinimalPokemon{
		ID:             p.ID,
		Name:           p.Name,
		Height:         float64(p.Height) / 10,  // decimetres to meters
		Weight:         float64(p.Weight) / 10,  // hectograms to kg
		BaseExperience: p.BaseExperience,
		Sprite:         p.Sprites.FrontDefault,
		Cries: map[string]string{
			"latest": p.Cries.Latest,
			"legacy": p.Cries.Legacy,
		},
	}
	
	// Extract types
	minimal.Types = make([]string, len(p.Types))
	for i, t := range p.Types {
		minimal.Types[i] = t.Type.Name
	}
	
	// Extract abilities
	minimal.Abilities = make([]AbilityInfo, len(p.Abilities))
	for i, a := range p.Abilities {
		minimal.Abilities[i] = AbilityInfo{
			Name:     a.Ability.Name,
			IsHidden: a.IsHidden,
		}
	}
	
	// Extract stats
	minimal.Stats = make(map[string]int)
	for _, s := range p.Stats {
		minimal.Stats[s.Stat.Name] = s.BaseStat
	}
	
	return minimal
}

func (c *Client) GetPokemonMinimal(name string) (MinimalPokemon, error) {
	var minimal MinimalPokemon
	
	// Get full data first
	fullPokemon, err := c.GetPokemonDetails(name)
	if err != nil {
		return minimal, err
	}
	
	// Transform to minimal
	minimal = fullPokemon.ToMinimal()
	return minimal, nil
}
