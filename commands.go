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
	baseUrl := "https://pokeapi.co/api/v2/location-area"
	if len(cfg.words) < 2 {
		return errors.New("You didn't pass the location")
	}

	locationUrl := baseUrl + "/" + cfg.words[1]

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