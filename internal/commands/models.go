package commands

import (
	"github.com/muhola/goPokedex/internal/cache"
	"github.com/muhola/goPokedex/internal/pokeapi"
)

type CliCommand struct {
	name        string
	description string
	callback    func(*Configuration, *cache.Cache, *pokeapi.Client, string, map[string]pokeapi.Pokemon) error
	config      *Configuration
}

type Configuration struct {
	Next     string
	Previous string
}
