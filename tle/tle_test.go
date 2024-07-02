package tle

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/chrishorton/spacerecon/config"
)

func TestGetTLE(t *testing.T) {
	config.LoadConfig(configDir)
	satelites, err := Get()
	assert.NotNil(t, satelites)
	assert.Nil(t, err)
}
