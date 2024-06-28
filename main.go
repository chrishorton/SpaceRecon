package main

import (
	"fmt"

	"github.com/manifoldco/promptui"

	"github.com/chrishorton/spacerecon/config"
	"github.com/chrishorton/spacerecon/spacerecon"
)

func main() {
	config.LoadConfig(".")
	prompt := promptui.Select{
		Label: "Select Command",
		Items: []string{"Query TLE", "Starlink Satellites", "Exit"},
	}
	for {
		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch result {
		case "Query TLE":
			// Handle TLE query
			fmt.Println("Querying TLE...")
			// Add TLE query functionality here

		case "Starlink Satellites":
			// Handle Starlink satellites query
			fmt.Println("Fetching Starlink satellites...")
			spacerecon.GetStarlinkSatellites()

		case "Exit":
			fmt.Println("Exiting...")
			return
		}
	}
}
