package main

import (
	"fmt"
	"os"
)

func commandHelp(c *config, parameter string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}

	return nil
}

func commandExit(c *config, parameter string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandMap(c *config, parameter string) error {
	locationsRes, err := c.pokeapiClient.GetLocations(c.Next)
	if err != nil {
		return err
	}

	c.Next = locationsRes.Next
	c.Previous = locationsRes.Previous

	for _, loc := range locationsRes.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(c *config, parameter string) error {
	if c.Previous == nil {
		return fmt.Errorf("you're on the first page")
	}
	locationsRes, err := c.pokeapiClient.GetLocations(c.Previous)
	if err != nil {
		return err
	}

	c.Next = locationsRes.Next
	c.Previous = locationsRes.Previous

	for _, loc := range locationsRes.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandExplore(c *config, parameter string) error {
	exploreRes, err := c.pokeapiClient.GetPokemon(parameter)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", exploreRes.Name)
	fmt.Println("Found Pokemon: ")

	for _, encounters := range exploreRes.PokemonEncounters {
		fmt.Printf(" - %s\n", encounters.Pokemon.Name)
	}

	return nil
}
