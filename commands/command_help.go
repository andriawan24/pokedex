package commands

import "fmt"

func help(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println()
	for _, command := range GetCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
