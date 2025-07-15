package main

import (
	"fmt"
	"github.com/lymvs/pokedexcli/internal/locationarea"
)

func commandHelp(_ *locationarea.Paginate) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
