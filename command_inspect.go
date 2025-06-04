package main

import (
	"fmt"
)

func commandInspect(cfg *config, pokemon string) error {
	fmt.Printf("Pokemon given: %s", pokemon)
	return nil
}
