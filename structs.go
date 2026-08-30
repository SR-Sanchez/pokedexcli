package main

type cliCommand struct {
	name        string
	description string
	callback     func(*config) error
}

type config struct {
	commands      map[string]cliCommand
	next          string
	previous      string
	pokeClient    Client
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type MapResult struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}