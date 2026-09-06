package main

import (
	"fmt"
	"os"
	"errors"
)

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for _, cmd := range cfg.commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config) error {
	baseUrl := "https://pokeapi.co/api/v2/location-area"
	url := cfg.next
	if url == "" {
		url = baseUrl
	}
	mapData, err := cfg.pokeClient.MakePokedexRequest(url)
	if err != nil {
		fmt.Println(err)
	}

	cfg.next = mapData.Next
	cfg.previous = mapData.Previous

	for _, location := range mapData.Results {
		fmt.Println(location.Name)
	}
	
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previous == ""{
		fmt.Println("you're on the first page")
		return nil
	}

	mapData, err := cfg.pokeClient.MakePokedexRequest(cfg.previous)
	if err != nil {
		fmt.Println(err)
	}

	cfg.next = mapData.Next
	cfg.previous = mapData.Previous

	for _, location := range mapData.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandExplore(cfg *config) error {
	locationUrl, err := GetUrl(cfg, "explore")
	if err != nil {
		return err
	}

	areaData, err := cfg.pokeClient.ListPokemonsInArea(locationUrl)
	if err != nil {
		return err
	}

	fmt.Println("Exploring " + areaData.Name + "...")
	for _, pokemonEncounter := range areaData.PokemonEncounters {
		fmt.Println("- " + pokemonEncounter.Pokemon.Name)
	}	
	
	return nil
}

func commandCatch(cfg *config) error {
	/* pokemonUrl, err := GetUrl(cfg, "catch")
	if err != nil {
		return err
	}
 */
 return nil
}

//helpers
func GetUrl(cfg *config, action string) (string, error) {
	baseUrl := "https://pokeapi.co/api/v2/"
	var route string
	var error string
	switch action {
		case "explore":
			route = "location-area/"
			error = "You didn't pass the location"
		case "catch":
		  route = "pokemon/"
		  error = "You didn't pass the pokemon name"
		default: 
		  route  = ""
			error = "Unknown action"
			return route, errors.New(error)
	}
	if len(cfg.words) < 2 {
		return route, errors.New(error)
	}

	searchUrl := baseUrl + route + cfg.words[1]
  return searchUrl, nil
}