package main

import (
	"fmt"
	"os"
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
	mapData := makePokedexRequest(cfg, url)

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

	mapData := makePokedexRequest(cfg, cfg.previous)

	for _, location := range mapData.Results {
		fmt.Println(location.Name)
	}

	return nil
}