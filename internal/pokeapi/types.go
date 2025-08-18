package pokeapi

// LocationResult represents a single item in the "results" array
type LocationResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// PokeAPIResponse represents the entire response from the PokeAPI
type PokeAPILocationResponse struct {
	Count    int              `json:"count"`
	Next     *string          `json:"next"`
	Previous *string          `json:"previous"`
	Results  []LocationResult `json:"results"`
}

type MinimalLocationArea struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon NamedAPIResource `json:"pokemon"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type StatData struct {
	StatName string `json:"name"`
}

type PokemonStats struct {
	Base_Stat int      `json:"base_stat"`
	Effort    int      `json:"effort"`
	Stat      StatData `json:"stat"`
}

type TypeInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonType struct {
	Slot int      `json:"slot"`
	Type TypeInfo `json:"type"`
}

type Pokemon struct {
	Name           string         `json:"name"`
	BaseExperiance int            `json:"base_experience"`
	Height         int            `json:"height"`
	Weight         int            `json:"weight"`
	Stats          []PokemonStats `json:"stats"`
	Types          []PokemonType  `json:"types"`
}

type Pokedex struct {
	Pokemons map[string]Pokemon
}
