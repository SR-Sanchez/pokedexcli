package pokedexClient

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

type PokemonAreaResult struct {
	Name                 string                     `json:"name"`
	PokemonEncounters    []PokemonEncounter         `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon        Pokemon          `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonInfo struct {
	Name              string  `json:"name"`
	BaseExperience    int     `json:"base_experience"`
}



