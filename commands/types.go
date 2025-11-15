package commands

import "github.com/andriawan24/pokedex/internal/pokeapi"

type cliCommand struct {
	name        string
	description string
	Callback    func(cfg *config) error
}

type config struct {
	client  pokeapi.Client
	nextURL *string
	prevURL *string
	pokedex map[string]pokeapi.Pokemon
	args    *string
}

func NewConfig(client pokeapi.Client) config {
	return config{
		client:  client,
		pokedex: make(map[string]pokeapi.Pokemon),
	}
}

func (cfg *config) SetArgs(args string) {
	cfg.args = &args
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex.",
			Callback:    exit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    help,
		},
		"map": {
			name:        "map",
			description: "Display 20 location areas and next page in the Pokemon World",
			Callback:    getMaps,
		},
		"mapb": {
			name:        "map",
			description: "Display previous 20 location areas in the Pokemon World",
			Callback:    getPrevMaps,
		},
		"explore": {
			name:        "explore [location area]",
			description: "Explore pokemons in an area, using area from map as a parameter",
			Callback:    explore,
		},
		"catch": {
			name:        "catch [pokemon]",
			description: "Throw a ball and try to catch a pokemon",
			Callback:    catch,
		},
		"inspect": {
			name:        "inspect [pokemon]",
			description: "See detail of pokemon from your pokedex",
			Callback:    inspect,
		},
	}
}
