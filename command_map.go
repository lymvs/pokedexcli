package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/lymvs/pokedexcli/internal/pokeapi"
)

func commandMap(c *pokeapi.Paginate) error {
	url := c.Next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	data, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return errors.New("Response failed")
	}
	if err != nil {
		return err
	}

	location := pokeapi.LocationArea{}
	if err := json.Unmarshal(data, &location); err != nil {
		return err
	}

	for _, result := range location.Results {
		fmt.Println(result.Name)
	}

	c.Next = location.Next
	c.Previous = location.Previous

	return nil
}
