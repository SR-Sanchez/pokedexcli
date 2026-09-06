package main

import (
	"bufio" // Provides buffered I/O, including Scanner for reading input.
	"fmt"
	"os"        // Provides access to operating-system functionality, including stdin.
	s "strings" // Provides functions for working with strings. "s" is an alias for the package.
)

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:     commandExit,
		},
		"help": {
			name:        "help",
			description: "Show help",
			callback:     commandHelp,
		},
		"map": {
			name:        "map",
			description: "displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "displays the names of 20 previous location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "After using the `map` command to find a location area, you can see a list of all the Pokémon located there",
			callback:    commandExplore,
		},
	}
}


// CleanInput takes a string of text and cleans it up.
//
// The text is converted to lowercase and then split into
// individual words using whitespace as the separator.
//
// For example:
// "Hello   WORLD" -> ["hello", "world"]
func CleanInput(text string) []string {
	// Convert the input to lowercase and split it into words.
	words := s.Fields(s.ToLower(text))

	// Return the cleaned list of words.
	return words
}

// startRepl waits for the user to enter a line of text
// in the terminal and returns that text as a string.
func startRepl(cfg *config) {
	// Create a Scanner that reads input from standard input
	// (the terminal/console).
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := CleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		cfg.words = words
		if cmd, ok := cfg.commands[words[0]]; ok {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}



