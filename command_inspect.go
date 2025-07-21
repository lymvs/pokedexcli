package main

import (
	"fmt"

	"github.com/lymvs/pokedexcli/internal/pokeapi"
)

func commandInspect(_ *pokeapi.Paginate, args []string) error {
	pokemon_name := args[0]

	if pokemon, found := Pokedex[pokemon_name]; !found {
		fmt.Println("you have not caught that pokemon")
	} else {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("-%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, pokemon_type := range pokemon.Types {
			fmt.Printf("-%s\n", pokemon_type.Type.Name)
		}
	}

	return nil
}
