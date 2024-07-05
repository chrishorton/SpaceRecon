package config

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	LoadConfig("../")
	if Cfg.Username == "" {
		t.Errorf("Username is empty")
	}
	if Cfg.Password == "" {
		t.Errorf("Password is empty")
	}
	if Cfg.CDMLimit == 0 {
		t.Errorf("CDMLimit is 0")
	}
}
