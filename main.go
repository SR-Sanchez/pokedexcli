package main

import (
	"github.com/SR-Sanchez/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokedexClient.NewClient(10)
	cmd := getCommands()
	pokemon := make(map[string]pokedexClient.PokemonInfo)
	startRepl(&config{ 
		commands: cmd,
		pokeClient: pokeClient,
		pokemon: pokemon,
	})
}
   