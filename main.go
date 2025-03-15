package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/muhola/goPokedex/internal/commands"
)

func main() {
	commands.InitComands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, "Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		commands.GetCommand(input)

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}
}
