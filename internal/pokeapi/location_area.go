package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocations(pageURL *string) (LocationArea, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error creating request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return LocationArea{}, fmt.Errorf("error fetching location area: %v", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error reading response: %v", err)
	}

	var la LocationArea
	if err := json.Unmarshal(data, &la); err != nil {
		return LocationArea{}, fmt.Errorf("error unmarshalling location area: %v", err)
	}

	return la, nil
}
