package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/muhola/goPokedex/internal/cache"
	"github.com/muhola/goPokedex/internal/commands"
	"github.com/muhola/goPokedex/internal/pokeapi"
)

const interval = 10 * time.Minute

func main() {
	pokedex := make(map[string]pokeapi.Pokemon)
	cache := cache.NewCache(interval)
	pokeClient := pokeapi.NewClient()
	commands.InitComands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, "Pokedex > ")
		//fmt.Println(os.Stderr, "Pokedex > ")   this fixed the Boot.Dev submit problem with Windows WSL
		if !scanner.Scan() {
			break
		}
		input := strings.Split(scanner.Text(), " ")
		parameter := ""

		if len(input) > 1 {
			parameter = input[1]
		}

		commands.GetCommand(input[0], parameter, cache, pokeClient, pokedex)

	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}
}
