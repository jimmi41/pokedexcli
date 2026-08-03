package pokeapi

type LocationAreaResponse struct{
	Count int
	Next *string
	Previous *string
	Results []Location
}

type Location struct{
	Name string
	Url string
}

type LocationAreaDetail struct {
	Name              string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonDetails struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	BaseExperience int            `json:"base_experience"`
	Height         int            `json:"height"`
	Weight         int            `json:"weight"`
	Abilities      []AbilitySlot  `json:"abilities"`
	Types          []TypeSlot     `json:"types"`
	Stats          []StatSlot     `json:"stats"`
}

type AbilitySlot struct {
	Ability  Ability `json:"ability"`
	IsHidden bool    `json:"is_hidden"`
	Slot     int     `json:"slot"`
}

type Ability struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type TypeSlot struct {
	Slot int         `json:"slot"`
	Type PokemonType `json:"type"`
}

type PokemonType struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type StatSlot struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}

type Stat struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}