package main

import (
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(text)
}

func main() {
	cleanInput("Hello, World!")
}
