package main

import (
	"fmt"
	"os"
)

func commandExit(conf *cliConfig, para string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
