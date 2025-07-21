package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/lymvs/pokedexcli/internal/pokeapi"
)

func commandExplore(c *pokeapi.Paginate, args []string) error {
	base_url := "https://pokeapi.co/api/v2/location-area"
	name := args[0]
	url := base_url + "/" + name

	var data []byte

	if cachedData, found := c.CacheResult.Get(url); found {
		data = cachedData
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return errors.New("Response failed")
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		c.CacheResult.Add(url, data)
	}

	locationName := pokeapi.LocationAreaName{}
	if err := json.Unmarshal(data, &locationName); err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, result := range locationName.PokemonEncounters {
		fmt.Printf("- %s\n", result.Pokemon.Name)
	}

	return nil
}
