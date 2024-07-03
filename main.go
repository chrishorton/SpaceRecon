package main

import (
    "github.com/chrishorton/spacerecon/cli"

	"github.com/chrishorton/spacerecon/config"
)

func main() {
    config.LoadConfig(".")
	cli.MainMenu()
}
