// Config test

package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("GOMON_SERVER_PORT", "8080")
	os.Setenv("GOMON_PROFILE_PORT", "6060")

	conf := Load()

	if conf.ServerPort != "8080" {
		t.Errorf("ServerPort is not 8080")
	}

	if conf.ProfilePort != "6060" {
		t.Errorf("ProfilePort is not 6060")
	}
}
