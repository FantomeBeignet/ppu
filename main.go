package main

import (
	"os"

	"git.sr.ht/~fantomebeignet/ppu/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
