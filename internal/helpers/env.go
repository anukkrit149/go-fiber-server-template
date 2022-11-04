package helpers

import (
	"os"
)

func GetEnv() (string, error) {
	// Fetch env for bootstrapping
	environ := os.Getenv("APP_MODE")
	if environ == "" {
		environ = "dev"
	}
	return environ, nil
}
