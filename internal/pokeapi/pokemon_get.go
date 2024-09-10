package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	if val, ok := c.cache.Get(url); ok {
		var p Pokemon
		if err := json.Unmarshal(val, &p); err != nil {
			return Pokemon{}, fmt.Errorf("error unmarshalling encounters: %v", err)
		}

		return p, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error creating request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("error fetching encounters: %v", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error reading response: %v", err)
	}

	var p Pokemon
	if err := json.Unmarshal(data, &p); err != nil {
		return Pokemon{}, fmt.Errorf("error unmarshalling encounters: %v", err)
	}

	c.cache.Add(url, data)
	return p, nil
}
