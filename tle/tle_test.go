package tle

import (
    "fmt"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/chrishorton/spacerecon/config"
)

func TestGetTLE(t *testing.T) {
	config.LoadConfig(configDir)
    // test with ISS
	satellites, err := Get(25544)

    fmt.Println(satellites)

	assert.NotNil(t, satellites)
	assert.Nil(t, err)
}
