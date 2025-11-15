package commands

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

func catch(cfg *config) error {
	if cfg.args == nil {
		return errors.New("pokemon name is empty")
	}

	name := strings.ToLower(*cfg.args)

	if _, ok := cfg.pokedex[name]; ok {
		fmt.Printf("%s is already on your pokedex\n", name)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := cfg.client.PokemonDetail(name)
	if err != nil {
		return err
	}

	chance := int(100 - float32(pokemon.BaseExperience)*0.42) // Mid level of base exp
	if chance < 15 {
		chance = 15
	} else if chance > 90 {
		chance = 90
	}

	actual := rand.Intn(100)

	if actual <= chance {
		fmt.Printf("%s was caught!\n", name)
		cfg.pokedex[name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}
