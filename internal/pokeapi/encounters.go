package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetEncounters(area string) (Encounters, error) {
	url := baseURL + "/location-area/" + area

	if val, ok := c.cache.Get(url); ok {
		var e Encounters
		if err := json.Unmarshal(val, &e); err != nil {
			return Encounters{}, fmt.Errorf("error unmarshalling encounters: %v", err)
		}

		return e, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Encounters{}, fmt.Errorf("error creating request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Encounters{}, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return Encounters{}, fmt.Errorf("error fetching encounters: %v", resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Encounters{}, fmt.Errorf("error reading response: %v", err)
	}

	var e Encounters
	if err := json.Unmarshal(data, &e); err != nil {
		return Encounters{}, fmt.Errorf("error unmarshalling encounters: %v", err)
	}

	c.cache.Add(url, data)
	return e, nil
}
