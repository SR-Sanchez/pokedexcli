package main

func main() {
	pokeClient := newClient(10)
	cmd := getCommands()
	startRepl(&config{ 
		commands: cmd,
		pokeClient: pokeClient,
	})
}
   