package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(location string) (LocationArea, error) {
	// store full url to make use of caching
	endpoint := "/location-area/" + location
	fullUrl := BaseURL + endpoint

	var locationArea LocationArea
	// cache hit
	data, ok := c.clientCache.Get(fullUrl)
	if ok {
		fmt.Println("cache hit!")
		if err := json.Unmarshal(data, &locationArea); err != nil {
			return locationArea, nil
		}

		return locationArea, nil
	}

	fmt.Println("cache miss!")
	// cache miss
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return locationArea, fmt.Errorf("error creating new request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationArea, fmt.Errorf("error making request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return locationArea, fmt.Errorf("bad status: %d %s",
			resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return locationArea, fmt.Errorf("Error reading body")
	}

	if err := json.Unmarshal(data, &locationArea); err != nil {
		return locationArea, nil
	}

	// cache the data
	c.clientCache.Set(fullUrl, data)

	return locationArea, nil
}

