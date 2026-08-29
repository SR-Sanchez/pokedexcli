package main

type cliCommand struct {
	name        string
	description string
	callback     func(*config) error
}

type config struct {
	commands map[string]cliCommand
}