package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/muhola/goPokedex/internal/cache"
	"github.com/muhola/goPokedex/internal/commands"
	"github.com/muhola/goPokedex/internal/pokeapi"
)

const interval = 10 * time.Minute

func main() {
	cache := cache.NewCache(interval)
	pokeClient := pokeapi.NewClient()
	commands.InitComands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, "Pokedex > ")
		//fmt.Println(os.Stderr, "Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		commands.GetCommand(input, cache, pokeClient)

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}
}
