package cdm

import (
	"fmt"
	"os"

	"github.com/chrishorton/spacerecon/config"
	"github.com/go-resty/resty/v2"
	flatbuffers "github.com/google/flatbuffers/go"
)

// Interface with the Space-Track API for public conjunctions
// Check first if the config has amount to fetch. if not, prompt user for amount

const (
	uriBase    = "https://www.space-track.org"
	authURL    = "/ajaxauth/login"
	requestURL = "/basicspacedata/query/class/cdm_public"
)

type Conjunction struct {
	Cdm_id               string `json:"cdm_id"`
	Created              string `json:"created"`
	Emergency_reportable string `json:"emergency_reportable"`
	Tca                  string `json:"tca"`
	Min_rng              string `json:"min_rng"`
	Pc                   string `json:"pc"`
	Sat_1_id             string `json:"sat_1_id"`
	Sat_1_name           string `json:"sat_1_name"`
	Sat1_object_type     string `json:"sat1_object_type"`
	Sat1_rcs             string `json:"sat1_rcs"`
	Sat_1_excl_vol       string `json:"sat_1_excl_vol"`
	Sat_2_id             string `json:"sat_2_id"`
	Sat_2_name           string `json:"sat_2_name"`
	Sat2_object_type     string `json:"sat2_object_type"`
	Sat2_rcs             string `json:"sat2_rcs"`
	Sat_2_excl_vol       string `json:"sat_2_excl_vol"`
}

func GetConjunctions() ([]Conjunction, error) {
	fmt.Println("Fetching Conjunctions...")
	var (
		username = config.Cfg.Username
		password = config.Cfg.Password
		limit    = config.Cfg.CDMLimit
	)

	client := resty.New()

	resp, err := client.R().
		SetFormData(map[string]string{
			"identity": username,
			"password": password,
		}).
		Post(uriBase + authURL)

	if err != nil || resp.StatusCode() != 200 {
		return nil, fmt.Errorf("error logging in: %v, %s", err, resp.Status())
	}

	fmt.Println("Fetching conjunctions... Limit is set to", limit)
	url := uriBase + requestURL + fmt.Sprintf("/limit/%d", limit)

	resp, err = client.R().
		SetResult(&[]Conjunction{}).
		Get(url)

	if err != nil {
		return nil, fmt.Errorf("error fetching conjunctions: %v", err)
	}

	PrintConjunctions(*resp.Result().(*[]Conjunction))

	return *resp.Result().(*[]Conjunction), nil
}

func PrintConjunctions(conjunctions []Conjunction) {
	// print conjunctions reversed so the most recent is first
	fmt.Println("Printing Conjunctions")
	for i := len(conjunctions) - 1; i >= 0; i-- {
		conjunctions[i].PrintConjunction()
	}
}

func (c Conjunction) PrintConjunction() {
	border := "====================================="
	fmt.Println()
	fmt.Println(border)
	fmt.Printf("Conjunction ID: %s\n", c.Cdm_id)
	fmt.Println(border)
	fmt.Printf("Created: %s\n", c.Created)
	fmt.Printf("TCA: %s\n", c.Tca)
	fmt.Printf("Minimum Range: %s\n", c.Min_rng)
	fmt.Printf("Probability of Collision: %s\n", c.Pc)
	fmt.Println(border)
	fmt.Printf("Satellite 1: %s (%s)\n", c.Sat_1_name, c.Sat_1_id)
	fmt.Printf("Satellite 2: %s (%s)\n", c.Sat_2_name, c.Sat_2_id)
	fmt.Println(border)
	fmt.Println()
}

func (c *Conjunction) SerializeConjunction() ([]byte, error) {
	builder := flatbuffers.NewBuilder(0)
	fmt.Println(c)
	return builder.Bytes, nil
}

func SerializeConjunctions(conjunctions []Conjunction) (bool, error) {
	// create a file for writing pass the filename as an arg
	fileName := "conjunctions.bin"
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return false, fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	for _, c := range conjunctions {
		bytes, err := c.SerializeConjunction()
		if err != nil {
			return false, fmt.Errorf("error serializing conjunction: %v", err)
		}
		// write bytes to file
		_, err = file.Write(bytes)
		if err != nil {
			return false, fmt.Errorf("error writing to file: %v", err)
		}
	}
	return true, nil
}
