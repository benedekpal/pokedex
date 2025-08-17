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
