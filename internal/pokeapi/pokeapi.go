package pokeapi

import (
	"encoding/json"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

func (c *Client) GetLocations(pageURL *string) (RespLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
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

	var respLocations RespLocations
	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&respLocations); err != nil {
		return RespLocations{}, err
	}

	return respLocations, nil
}
