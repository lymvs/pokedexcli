package main

import "strings"

func cleanInput(text string) []string {
	var words []string

	words = strings.Fields(strings.ToLower(text))

	return words
}
