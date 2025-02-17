package main

//CLI Commands

import (
	"fmt"
	"math/rand"
	"os"
)

// Help - Displays command name and descriptions
func commandHelp(c *config, parameter string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}

	return nil
}

// Exit - Exits CLI
func commandExit(c *config, parameter string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// Map - Displays Next 20 Locations
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

// Mapb - Displays previous 20 Locations
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

// Explore - Displays list of pokemon in a given location
func commandExplore(c *config, parameter string) error {
	exploreRes, err := c.pokeapiClient.GetPokemonInLocation(parameter)
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

// Catch - Attempts to catch a pokemon and add it to pokedex
func commandCatch(c *config, name string) error {
	pokemon, err := c.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	if rand.Intn(pokemon.BaseExperience) <= 50 {
		fmt.Printf("%s was caught!\n", name)
		c.Pokedex[name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", name)
		catchRate := 50.0 / float64(pokemon.BaseExperience) * 100
		fmt.Printf("catch rate: %.2f%%\n", catchRate)
	}

	return nil
}

// Inspect - Displays info on a caught pokemon
func commandInspect(c *config, name string) error {
	if name == "" {
		return fmt.Errorf("inspect requires a pokemon name")
	}

	pokedata, ok := c.Pokedex[name]

	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokedata.Name)
	fmt.Printf("Height: %d\n", pokedata.Height)
	fmt.Printf("Weight: %d\n", pokedata.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokedata.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, poketype := range pokedata.Types {
		fmt.Printf("  - %s\n", poketype.Type.Name)
	}

	return nil
}

// Pokedex - Displays all caught pokemon
func commandPokedex(c *config, parameter string) error {
	if len(c.Pokedex) == 0 {
		fmt.Println("You have not caught any pokemon")
		return nil
	}
	for name := range c.Pokedex {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
