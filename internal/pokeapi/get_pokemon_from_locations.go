package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type pokemonEncounters struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c Client) GetPokemonFromLocation(location string) (pokemonEncounters, error) {
	url := baseURL + "location-area/" + location + "/"

	if val, ok := c.cache.Get(url); ok {
		var encounters pokemonEncounters
		err := json.Unmarshal(val, &encounters)
		if err != nil {
			return pokemonEncounters{}, err
		}
		return encounters, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokemonEncounters{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return pokemonEncounters{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return pokemonEncounters{}, nil
	}

	var encounters pokemonEncounters
	err = json.Unmarshal(body, &encounters)
	if err != nil {
		return pokemonEncounters{}, err
	}

	c.cache.Add(url, body)

	return encounters, err
}
