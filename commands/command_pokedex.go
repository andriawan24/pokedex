package commands

import (
	"errors"
	"fmt"
)

func pokedex(cfg *config) error {
	if len(cfg.pokedex) == 0 {
		return errors.New("You haven't caught any pokemon yet")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.pokedex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
