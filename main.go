package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/jimmi41/pokedexcli/internal/pokeapi"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	commands := getCommands()
	
	cfg := &config{
    	pokedex: make(map[string]pokeapi.PokemonDetails),
	}

	for {
		fmt.Print("Pokedex > ")

		scanned := scanner.Scan()
		if !scanned {
			break
		}

		input := scanner.Text()
		words := cleanInput(input)

		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := commands[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(cfg, words[1:])
		if err != nil {
			fmt.Println(err)
		}
	}
}
