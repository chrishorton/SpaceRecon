package cdm

import (
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
