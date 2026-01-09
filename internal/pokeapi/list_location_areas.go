package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) ListLocationAreas() (LocationAreasResp, error) {
	endpoint := "/location-area"
	url := BaseURL + endpoint
	var locationAreaResp LocationAreasResp

	req, err := http.NewRequest("GET", url, nil)
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

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&locationAreaResp); err != nil {
		return locationAreaResp, fmt.Errorf("error decoding json: %v", err)
	}

	return locationAreaResp, nil
}
