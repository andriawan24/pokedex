package commands

import (
	"errors"
	"fmt"
)

func explore(cfg *config) error {
	if cfg.args == nil {
		return errors.New("Input is empty")
	}

	locationArea := *cfg.args

	fmt.Printf("Exploring %s...\n", locationArea)

	locationDetailResponse, err := cfg.client.ListLocationDetail(locationArea)
	if err != nil {
		return err
	}

	fmt.Println("Found pokemon:")

	for _, detail := range locationDetailResponse.PokemonEncounters {
		fmt.Printf("- %s\n", detail.Pokemon.Name)
	}

	return nil
}
