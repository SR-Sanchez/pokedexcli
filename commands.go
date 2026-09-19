package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"
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
	pokemonUrl, err := GetUrl(cfg, "catch")
	if err != nil {
		return err
	}
	pokemonData, err := cfg.pokeClient.CatchPokemon(pokemonUrl); 
	if err != nil {
		return err
	}

	catchingChance := 1.0 - float64(pokemonData.BaseExperience)/300.0
	
	fmt.Println("Throwing a Pokeball at " + pokemonData.Name + "...")
	time.Sleep(time.Second * 3)
	if rand.Float64() < catchingChance {
		cfg.pokemon[pokemonData.Name] = pokemonData
		fmt.Printf("%s was caught!\nYou may now inspect it with the inspect command.\n", pokemonData.Name)
		return nil
	}
	fmt.Printf("%v got away...\n", pokemonData.Name)
	return nil
}

func commandInspect(cfg *config) error {
	if len(cfg.words) < 2 {
		fmt.Println("No pokemon name was passed")
		return nil
	}
	pokemonName := cfg.words[1]
	pokemon, ok := cfg.pokemon[pokemonName]
  if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n",  pokemon.Name)
	fmt.Printf("Height: %v\n",  pokemon.Height)
	fmt.Printf("Weight: %v\n",  pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("%s: %v\n", stat.Stat.Name, stat.BaseStat )
	}
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("-%s\n", pokemonType.Type.Name)
	}
	
	return nil
}

func commandPokedex(cfg *config) error {
	if len(cfg.pokemon) == 0 {
		fmt.Println("You haven't catch any pokemon yet")
		return nil
	}
	fmt.Println("Your pokedex:")
	for _, value := range cfg.pokemon{
		fmt.Printf("- %s\n", value.Name)
	}
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