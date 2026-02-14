package main

import "strings"

func cleanInput(text string) []string {
	blow := strings.Fields(text)
	for i := range len(blow) {
		blow[i] = strings.ToLower(blow[i])
	}
	return blow
}
