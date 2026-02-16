package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/RagnaCron/pokedexcli/internal/pokeapi"
	"github.com/RagnaCron/pokedexcli/internal/pokecache"
	"github.com/RagnaCron/pokedexcli/internal/pokecommand"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	conf := pokecommand.Config{
		Cache:   pokecache.NewCache(10 * time.Minute),
		Pokedex: make(map[string]pokeapi.Pokemon),
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, ok := pokecommand.Get()[commandName]
		if ok {
			if err := command.Callback(&conf, args...); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
