package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "pokemon/" + name + "/"

	if val, ok := c.cache.Get(url); ok {
		var pokemonRes Pokemon
		err := json.Unmarshal(val, &pokemonRes)
		if err != nil {
			return Pokemon{}, nil
		}
		return pokemonRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, nil
	}

	var pokemonRes Pokemon
	err = json.Unmarshal(body, &pokemonRes)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, body)

	return pokemonRes, nil
}
