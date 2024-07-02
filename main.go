package main

import (
	"fmt"

	"github.com/manifoldco/promptui"

	"github.com/chrishorton/spacerecon/config"
	"github.com/chrishorton/spacerecon/tle"
    "github.com/chrishorton/spacerecon/starlink"
)

func main() {
	config.LoadConfig(".")
	prompt := promptui.Select{
		Label: "Select Command",
		Items: []string{"TLE and OMM", "Starlink Satellites", "Exit"},
	}

    TLESubmenu := promptui.Select{
        Label: "Select TLE Command",
        Items: []string{"Query TLE", "Load TLE From File", "Load TLE From Space-Track.org", "Main Menu"},
    }

	for {
		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch result {
		case "Query TLE":
            _, result, err := TLESubmenu.Run()
            if err != nil {
                fmt.Printf("Prompt failed %v\n", err)
                return
            }

            for {
                switch result {
                case "Query TLE":
                    fmt.Println("Fetching TLE...")
                    tles, err := tle.Get()
                    if err != nil {
                        fmt.Println("Error fetching TLE: ", err)
                    }
                    for _, t := range tles {
                        fmt.Println(t)
                    }
                case "Load TLE From File":
                    fmt.Println("Loading TLE from file...")
                case "Load TLE From Space-Track.org":
                    fmt.Println("Loading TLE from Space-Track.org...")
                case "Exit":
                    return
                }
            }

		case "Starlink Satellites":
			// Handle Starlink satellites query
			fmt.Println("Fetching Starlink satellites...")
            sats, err := starlink.Get()
            if err != nil {
                fmt.Println("Error fetching Starlink satellites: ", err)
            }
            for _, s := range sats {
                fmt.Println(s)
            }

		case "Exit":
			fmt.Println("Exiting...")
			return
		}
	}
}
