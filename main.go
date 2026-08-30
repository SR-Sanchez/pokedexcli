package main

func main() {
	cmd := getCommands()
	startRepl(&config{ commands: cmd})
}
   