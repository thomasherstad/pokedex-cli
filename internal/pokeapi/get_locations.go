package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"

	"log"

	"net/http"
)

type mapResults struct {
	Count     int     `json:"count"`
	Next      *string `json:"next"`
	Previous  *string `json:"previous"`
	Locations []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func (c Client) GetLocations(pageURL *string) (mapResults, error) {
	url := baseURL + "/location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationResponse := mapResults{}
		err := json.Unmarshal(val, &locationResponse)
		if err != nil {
			return mapResults{}, err
		}
		return locationResponse, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return mapResults{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with error code %d\n and body %s\n", resp.StatusCode, resp.Body)
	}
	if err != nil {
		fmt.Println(err)
		return mapResults{}, err
	}

	var locations mapResults
	err = json.Unmarshal(body, &locations)
	if err != nil {
		fmt.Println(err)
	}

	c.cache.Add(url, body)

	return locations, nil
}
