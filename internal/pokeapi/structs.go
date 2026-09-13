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

type PokemonStat struct {
	Name           string     `json:"name"` 
}

type PokemonStats struct {
	BaseStat        int            `json:"base_stat"`
	Stat            PokemonStat    `json:"stat"`
}

type PokemonTypes struct {
	Type struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"type"`
}

type PokemonInfo struct {
	Name              string  `json:"name"`
	BaseExperience    int     `json:"base_experience"`
	Height            int     `json:"height"`
	Weight            int     `json:"weight"`
	Stats             []PokemonStats  `json:"stats"`
	Types             []PokemonTypes   `json:"types"`
}



