package pokeapi

//Pokeapi Requests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

// Get a list of locations
func (c *Client) GetLocations(pageURL *string) (RespLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		respLocations := RespLocations{}
		err := json.Unmarshal(val, &respLocations)
		if err != nil {
			return RespLocations{}, err
		}
		return respLocations, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocations{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocations{}, err
	}

	var respLocations RespLocations
	err = json.Unmarshal(data, &respLocations)
	if err != nil {
		return RespLocations{}, err
	}

	c.cache.Add(url, data)
	return respLocations, nil
}

// Get data on a single location
func (c *Client) GetPokemonInLocation(location string) (RespSingleLocation, error) {
	url := baseURL + "/location-area/" + location
	if location == "" {
		return RespSingleLocation{}, fmt.Errorf("explore requires a location parameter")
	}

	if val, ok := c.cache.Get(url); ok {
		respSingleLocation := RespSingleLocation{}
		err := json.Unmarshal(val, &respSingleLocation)
		if err != nil {
			return RespSingleLocation{}, err
		}
		return respSingleLocation, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespSingleLocation{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespSingleLocation{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespSingleLocation{}, err
	}

	var respSingleLocation RespSingleLocation
	err = json.Unmarshal(data, &respSingleLocation)
	if err != nil {
		return RespSingleLocation{}, err
	}

	c.cache.Add(url, data)
	return respSingleLocation, nil
}

// Get data on a pokemon
func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name
	if name == "" {
		return Pokemon{}, fmt.Errorf("catch requires a pokemon name")
	}

	if val, ok := c.cache.Get(url); ok {
		pokemon := Pokemon{}
		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var pokemon Pokemon
	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)
	return pokemon, nil
}
