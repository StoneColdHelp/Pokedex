package main

import (
	"strings"
)

func main() {
	cleanInput("HelLo, WOrld!")
}

func cleanInput(text string) []string {
	text = strings.TrimSpace(strings.ToLower(text))
	if text == "" {
		return []string{}
	}
	return strings.Fields(text)
}
