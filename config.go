package main

import (
	"github.com/jimmi41/pokedexcli/internal/pokeapi"
)

type config struct {
	next     *string
	previous *string
	pokedex map[string]pokeapi.PokemonDetails
}
