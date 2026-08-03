package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jimmi41/pokedexcli/internal/pokecache"
)

const (
	baseURL = "https://pokeapi.co/api/v2/location-area"
)

var cache = pokecache.NewCache(5 * time.Minute)

func GetLocationAreas(url *string) (LocationAreaResponse, error) {
	// STEP 1:
	// Decide which URL to call.
	fullURL := baseURL
	if url != nil {
		fullURL = *url
	}

	// STEP 2:
	// Make HTTP GET request
	if data, found := cache.Get(fullURL); found {
		var locationAreaResp LocationAreaResponse
		if err := json.Unmarshal(data, &locationAreaResp); err != nil {
			return LocationAreaResponse{}, fmt.Errorf("failed to decode cached JSON: %w", err)
		}
		return locationAreaResp, nil
	}

	// STEP 3:
	// Make HTTP GET request.
	res, err := http.Get(fullURL)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("network error fetching locations: %w", err)
	}
	defer res.Body.Close()


	// STEP 5:
	// Check HTTP status code
	// Return an error if status isn't OK
	if res.StatusCode > 299 {
		return LocationAreaResponse{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	// STEP 6:
	// Decode JSON into LocationAreaResponse
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaResponse{}, fmt.Errorf("failed to read response body: %w", err)
	}

	cache.Add(fullURL, data)

	var locationAreaResp LocationAreaResponse
	if err := json.Unmarshal(data, &locationAreaResp); err != nil {
		return LocationAreaResponse{}, fmt.Errorf("failed to decode response JSON: %w", err)
	}

	// STEP 7:
	// Return decoded response
	return locationAreaResp, nil
}

func GetLocationDetail(url *string) (LocationAreaDetail, error) {
	// STEP 1:
	fullURL := ""
	if url == nil {
		return LocationAreaDetail{}, fmt.Errorf("url cannot be nil")
	}
	if url != nil {
		fullURL = *url
	}

	// STEP 2:
	// Make HTTP GET request
	if data, found := cache.Get(fullURL); found {
		var locationAreaDetail LocationAreaDetail
		if err := json.Unmarshal(data, &locationAreaDetail); err != nil {
			return LocationAreaDetail{}, fmt.Errorf("failed to decode cached JSON: %w", err)
		}
		return locationAreaDetail, nil
	}

	// STEP 3:
	// Make HTTP GET request.
	res, err := http.Get(fullURL)
	if err != nil {
		return LocationAreaDetail{}, fmt.Errorf("network error fetching locations: %w", err)
	}
	defer res.Body.Close()


	// STEP 5:
	// Check HTTP status code
	// Return an error if status isn't OK
	if res.StatusCode > 299 {
		return LocationAreaDetail{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	// STEP 6:
	// Decode JSON into LocationAreaDetail
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreaDetail{}, fmt.Errorf("failed to read response body: %w", err)
	}

	cache.Add(fullURL, data)

	var locationAreaDetail LocationAreaDetail
	if err := json.Unmarshal(data, &locationAreaDetail); err != nil {
		return LocationAreaDetail{}, fmt.Errorf("failed to decode response JSON: %w", err)
	}

	// STEP 7:
	// Return decoded response
	return locationAreaDetail, nil
}

func GetPokemonDetails(url *string) (PokemonDetails, error) {
	// STEP 1:
	fullURL := ""
	if url == nil {
		return PokemonDetails{}, fmt.Errorf("url cannot be nil")
	}
	if url != nil {
		fullURL = *url
	}

	// STEP 2:
	// Make HTTP GET request
	if data, found := cache.Get(fullURL); found {
		var pokemonDetails PokemonDetails
		if err := json.Unmarshal(data, &pokemonDetails); err != nil {
			return PokemonDetails{}, fmt.Errorf("failed to decode cached JSON: %w", err)
		}
		return pokemonDetails, nil
	}

	// STEP 3:
	// Make HTTP GET request.
	res, err := http.Get(fullURL)
	if err != nil {
		return PokemonDetails{}, fmt.Errorf("network error fetching pokemon: %w", err)
	}
	defer res.Body.Close()


	// STEP 5:
	// Check HTTP status code
	// Return an error if status isn't OK
	if res.StatusCode > 299 {
		return PokemonDetails{}, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	// STEP 6:
	// Decode JSON into PokemonDetails
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonDetails{}, fmt.Errorf("failed to read response body: %w", err)
	}

	cache.Add(fullURL, data)

	var pokemonDetails PokemonDetails
	if err := json.Unmarshal(data, &pokemonDetails); err != nil {
		return PokemonDetails{}, fmt.Errorf("failed to decode response JSON: %w", err)
	}

	// STEP 7:
	// Return decoded response
	return pokemonDetails, nil
}