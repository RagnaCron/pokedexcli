package main

import (
	"errors"
	"fmt"
)

func commandExplore(conf *cliConfig, para string) error {
	if para == "" {
		return errors.New("missing secound parameter")
	}
	fmt.Printf("Exploring %v...\n", para)

	return nil
}
