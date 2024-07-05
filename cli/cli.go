package cli

import (
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"

	"github.com/chrishorton/spacerecon/cdm"
	"github.com/chrishorton/spacerecon/tle"
)

func MainMenu() {
	prompt := promptui.Select{
		Label: "Select an option",
		Items: []string{"TLE", "Conjunctions", "Exit"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case "TLE":
		tleMenu()
	case "Conjunctions":
		fmt.Println("Conjunctions selected.")
		conjunctionsMenu()
	case "Exit":
		return
	default:
		fmt.Println("Invalid choice. Please try again.")
		MainMenu()
	}
}

func conjunctionsMenu() {
	fmt.Println("Conjunctions Menu")
	prompt := promptui.Select{
		Label: "Conjunctions Options",
		Items: []string{"Load from File", "Load from Space-Track.org", "Back"},
	}
	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case "Load from File":
		fmt.Println("Load from File selected.")
	case "Load from Space-Track.org":
		// no more input needed
		cdm.GetConjunctions()
	case "Back":
		MainMenu()
	default:
		fmt.Println("Invalid choice. Please try again.")
		conjunctionsMenu()
	}
}

func tleMenu() {
	prompt := promptui.Select{
		Label: "TLE Options",
		Items: []string{"Load from File", "Load from Space-Track.org", "Back"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch result {
	case "Load from File":
		loadFromFile()
		tleMenu()
	case "Load from Space-Track.org":
		loadTLEFromSpaceTrack()
		tleMenu()
	case "Back":
		MainMenu()

	default:
		fmt.Println("Invalid choice. Please try again.")
		tleMenu()
	}
}

func loadFromFile() {
	prompt := promptui.Prompt{
		Label: "Enter the filename, or 'exit' to return to the previous menu",
	}

	filename, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if filename == "" {
		fmt.Println("Invalid filename. Please try again.")
		loadFromFile()
		return
	} else if filename == "exit" {
		tleMenu()
	}

	fmt.Printf("Loading from file: %s\n", filename)
	// Add file loading logic here

}

func loadTLEFromSpaceTrack() {
	for {
		prompt := promptui.Prompt{
			Label: "Enter the NORAD_ID of the satellite to load, or 'exit' to return to the previous menu",
		}

		noradIDStr, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		if noradIDStr == "" {
			fmt.Println("Invalid NORAD_ID. Please try again.")
			continue
		} else if noradIDStr == "exit" {
			tleMenu()
		}

		noradID, err := strconv.Atoi(noradIDStr)
		if err != nil {
			fmt.Println("Invalid NORAD_ID. Please enter a valid integer.")
			continue
		}

		fmt.Printf("Loading from Space-Track.org for NORAD_ID: %d\n", noradID)
		tleRes, err := tle.Get(noradID)
		if err != nil {
			fmt.Printf("Error loading TLE: %v\n", err)
			continue
		}

		tle.PrintTLEs(tleRes)

		break
	}
}
