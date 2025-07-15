package locationarea

import (
	"encoding/json"
	"net/http"
	"fmt"
	"io"
	"errors"
)

func CommandMap(c *Paginate) error {
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

	location := LocationArea{}
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