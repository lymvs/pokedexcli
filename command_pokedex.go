package main

import (
	"fmt"

	"github.com/lymvs/pokedexcli/internal/pokeapi"
)

func commandPokedex(_ *pokeapi.Paginate, _ []string) error {
	if len(Pokedex) == 0 {
		fmt.Println("Your Pokedex is empty")
	} else {
		fmt.Println("Your Pokedex:")
		for _, pokemon := range Pokedex {
			fmt.Printf("- %s\n", pokemon.Name)
		}
	}

	return nil
}
