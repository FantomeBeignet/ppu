package main

import (
	"log"

	"git.sr.ht/~fantomebeignet/ppu/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("error: %s\n", err.Error())
	}
}
