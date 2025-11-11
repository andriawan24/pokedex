package main

import (
	"strings"
)

func cleanInput(text string) []string {
	var result []string

	trimText := strings.TrimSpace(text)
	if trimText == "" {
		return result
	}

	splitText := strings.Split(trimText, " ")

	for _, item := range splitText {
		if item != "" {
			result = append(result, strings.ToLower(item))
		}
	}

	return result
}
