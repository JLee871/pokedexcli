package pokeapi

type Location struct {
	Name string
	URL  string
}

type RespLocations struct {
	Count    int        `json:"count"`
	Next     *string    `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}
