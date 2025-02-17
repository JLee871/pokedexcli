package main

import (
	"time"

	"github.com/JLee871/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(10*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
		Pokedex:       make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
