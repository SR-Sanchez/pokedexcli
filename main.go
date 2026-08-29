package main

func main() {
	cmd := getCommands()
	startRepl(&config{cmd})
}
   