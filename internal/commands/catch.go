package commands

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/muhola/goPokedex/internal/cache"
	"github.com/muhola/goPokedex/internal/pokeapi"
)

func commandCatch(configuration *Configuration, pokeCache *cache.Cache, pokeClient *pokeapi.Client, pokemon string, pokedex map[string]pokeapi.Pokemon) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	findPokemon := pokeClient.GetPokemon(pokemon, pokeCache)

	// Create a random number generator at program start
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	maxBaseExp := 255.0
	catchChance := 1.0 - (float64(findPokemon.BaseExperience) / maxBaseExp)

	// Ensure catch chance is at least 10% and at most 90%
	if catchChance < 0.1 {
		catchChance = 0.1
	} else if catchChance > 0.9 {
		catchChance = 0.9
	}

	rollToCatchPokemon := rng.Float64()
	caughtPokemon := rollToCatchPokemon <= catchChance

	fmt.Printf("%.1f%% change to catch %s\n", catchChance*100, findPokemon.Name)
	if caughtPokemon {
		fmt.Printf("Congratulations! You caught %s!\n", findPokemon.Name)
		pokedex[findPokemon.Name] = findPokemon
	} else {
		fmt.Printf("Oh no! %s broke free and fled!\n", findPokemon.Name)
	}
	return nil
}
