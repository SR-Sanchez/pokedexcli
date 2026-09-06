package main

import (
	"github.com/SR-Sanchez/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokedexClient.NewClient(10)
	cmd := getCommands()
	startRepl(&config{ 
		commands: cmd,
		pokeClient: pokeClient,
	})
}
   