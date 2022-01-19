package main

import (
	"os"

	"github.com/silvanocerza/joyous-json/cmd"
)

func main() {
	if err := cmd.Root().Execute(); err != nil {
		os.Exit(1)
	}
}
