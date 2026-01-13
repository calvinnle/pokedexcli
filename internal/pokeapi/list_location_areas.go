package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResp, error) {
	// store full url to make use of caching
	endpoint := "/location-area?offset=0&limit=20"
	fullUrl := BaseURL + endpoint
	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	var locationAreaResp LocationAreasResp

	// cache hit
	data, ok := c.clientCache.Get(fullUrl)
	if ok {
		fmt.Println("cache hit!")
		if err := json.Unmarshal(data, &locationAreaResp); err != nil {
			return locationAreaResp, nil
		}

		return locationAreaResp, nil
	}

	fmt.Println("cache miss!")
	// cache miss
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return locationAreaResp, fmt.Errorf("error creating new request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreaResp, fmt.Errorf("error making request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return locationAreaResp, fmt.Errorf("bad status: %d %s",
			resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return locationAreaResp, fmt.Errorf("Error reading body")
	}

	if err := json.Unmarshal(data, &locationAreaResp); err != nil {
		return locationAreaResp, nil
	}

	// cache the data
	c.clientCache.Set(fullUrl, data)

	return locationAreaResp, nil
}
