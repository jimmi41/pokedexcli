package main

import (
	"fmt"
	"os"
	"github.com/jimmi41/pokedexcli/internal/pokeapi"
    "math/rand"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas",
			callback:     commandMap,
		},
		"explore": {
			name:        "explore",
			description: "Finds the next 20 location areas and displays them",
			callback:     commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a Pokémon",
			callback:     commandCatch,
		},
        "pokedex": {
			name:        "pokedex",
			description: "Displays the Pokémon you have caught",
			callback:     commandPokedex,
		},
	}
}

func commandHelp(cfg *config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	commands := getCommands()

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
func commandExit(cfg *config, args []string) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return nil
}

func commandMap(cfg *config, args []string) error {
    // Step 1: Decide which URL to request
    var url string
    if cfg.next == nil {
        url = "https://pokeapi.co/api/v2/location-area"
    } else {
        url = *cfg.next
    }

    // Step 2: Ask the API package for data
    response, err := pokeapi.GetLocationAreas(&url)
    if err != nil {
        return err
    }

    // Step 3: Save pagination information
    cfg.next = response.Next
    cfg.previous = response.Previous

    // Step 4: Print all location names
    for _, location := range response.Results {
        fmt.Println(location.Name)
    }

    return nil
}

func commandExplore(cfg *config, args []string) error {
    if len(args) == 0 {
        return fmt.Errorf("no location specified")
    }
    loc := args[0]

    // Step 1: Decide which URL to request
    var url string
	url = "https://pokeapi.co/api/v2/location-area/"+loc

    // Step 2: Ask the API package for data
    response, err := pokeapi.GetLocationDetail(&url)
    if err != nil {
        return err
    }

    // Step 3: Print all location names
    for _, pokemon := range response.PokemonEncounters {
        fmt.Println(pokemon.Pokemon.Name)
    }

    return nil
}

func commandCatch(cfg *config, args []string) error {
    if len(args) == 0 {
        return fmt.Errorf("no Pokémon specified")
    }
    pokemon := args[0]

    // Step 1: Decide which URL to request
    var url string
    url = "https://pokeapi.co/api/v2/pokemon/" + pokemon

    // Step 2: Ask the API package for data
    response, err := pokeapi.GetPokemonDetails(&url)
    if err != nil {
        return err
    }
    fmt.Printf("Throwing a Pokeball at %s...\n", response.Name)

    if(tryCatch(response.BaseExperience)) {
        cfg.pokedex[response.Name] = response
        fmt.Printf("You caught %s!\n", response.Name)
    } else {
        fmt.Printf("%s escaped!\n", response.Name)
    }
    
    return nil
}

func tryCatch(baseExperience int) bool {
	// Minimum 10% chance, maximum 90% chance
	catchChance := 100 - (baseExperience / 2)

	if catchChance < 10 {
		catchChance = 10
	}
	if catchChance > 90 {
		catchChance = 90
	}

	roll := rand.Intn(100) // 0-99

	return roll < catchChance
}

func commandPokedex(cfg *config, args []string) error {
    if len(cfg.pokedex) == 0 {
        fmt.Println("You haven't caught any Pokémon yet!")
        return nil
    }

    fmt.Println("Your Pokedex:")
    for name := range cfg.pokedex {
        fmt.Println("- " + name)
    }
    return nil
}
