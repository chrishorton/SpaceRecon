package spacerecon

import (
	"testing"

	"github.com/chrishorton/spacerecon/config"
	"github.com/stretchr/testify/assert"
)

const configDir = ".."

func TestGetStarlinkSatellites(t *testing.T) {
	config.LoadConfig(configDir)
	satelites, err := GetStarlinkSatellites()
	assert.NotNil(t, satelites)
	assert.Nil(t, err)
}

func TestGetTLE(t *testing.T) {
	config.LoadConfig(configDir)
	satelites, err := GetTLE()
	assert.NotNil(t, satelites)
	assert.Nil(t, err)
}
