package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/andriawan24/pokedex/commands"
	"github.com/andriawan24/pokedex/internal/pokeapi"
)

const (
	timeout       = 30 * time.Second
	cacheDuration = 10 * time.Second
)

func startRepl() {
	httpClient := pokeapi.NewClient(timeout, cacheDuration)
	cfg := commands.NewConfig(httpClient)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		text := scanner.Text()
		texts := cleanInput(text)
		if len(texts) == 0 {
			fmt.Println("Command is empty")
			continue
		}

		commandText := texts[0]
		command, exists := commands.GetCommands()[commandText]
		if exists {
			if len(texts) > 1 {
				args := strings.Join(texts[1:], "")
				cfg.SetArgs(args)
			}

			err := command.Callback(&cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Command not found")
			continue
		}
	}
}

func cleanInput(text string) []string {
	var result []string

	trimText := strings.TrimSpace(text)
	if trimText == "" {
		return result
	}

	splitText := strings.SplitSeq(trimText, " ")
	for item := range splitText {
		if item != "" {
			result = append(result, strings.ToLower(item))
		}
	}

	return result
}
