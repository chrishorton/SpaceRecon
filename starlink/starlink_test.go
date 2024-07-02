package starlink

import (
	"testing"

	"github.com/chrishorton/spacerecon/config"
	"github.com/stretchr/testify/assert"
)

func TestGetStarlinkSatellites(t *testing.T) {
	config.LoadConfig(configDir)
	satelites, err := Get()
	assert.NotNil(t, satelites)
	assert.Nil(t, err)
}

