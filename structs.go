package main

import (
	"github.com/SR-Sanchez/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback     func(*config) error
}

type config struct {
	commands      map[string]cliCommand
	next          string
	previous      string
	pokeClient    pokedexClient.Client
	words         []string
	pokemon       map[string]pokedexClient.PokemonInfo
}

