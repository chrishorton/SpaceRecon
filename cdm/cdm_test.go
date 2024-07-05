package cdm

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/chrishorton/spacerecon/config"
)

func TestGetConjunctions(t *testing.T) {
	config.LoadConfig("../")
	_, err := GetConjunctions()
	if err != nil {
		t.Errorf("Error fetching conjunctions: %v", err)
	}
}

func loadTestData() ([]Conjunction, error) {
	filePath := "testData/data.json"
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var conjunctions []Conjunction
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conjunctions)
	if err != nil {
		return nil, fmt.Errorf("error decoding JSON: %v", err)
	}
	return conjunctions, nil
}

func TestSerialize(t *testing.T) {
	config.LoadConfig("../")
	data, err := loadTestData()
	if err != nil {
		t.Errorf("Error loading test data: %v", err)
	}
	_, err = SerializeConjunctions(data)
	if err != nil {
		t.Errorf("Error serializing conjunctions: %v", err)
	}
}
