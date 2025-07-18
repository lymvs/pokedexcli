package main

import (
	"fmt"

	"github.com/lymvs/pokedexcli/internal/pokeapi"
)

func commandHelp(_ *pokeapi.Paginate) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
