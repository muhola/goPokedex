package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	commands.initComands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, "Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		clicmd, ok := commands.cliComands[strings.ToLower(input)]
		if ok {
			cmdError := clicmd.callback()
			if cmdError != nil {
				fmt.Errorf(cmdError.Error())
			}

		} else {
			fmt.Println("Unknown command")
		}

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}
}
