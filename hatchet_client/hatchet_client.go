package hatchet_client

import (
	"errors"
	"os"

	v1 "github.com/hatchet-dev/hatchet/pkg/v1"
	"github.com/joho/godotenv"
)

func HatchetClient() (v1.HatchetClient, error) {
	if _, err := os.Stat(".env"); os.IsExist(err) {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	// check for HATCHET_CLIENT_TOKEN
	token := os.Getenv("HATCHET_CLIENT_TOKEN")
	if token == "" {
		return nil, errors.New("HATCHET_CLIENT_TOKEN is not set")
	}

	return v1.NewHatchetClient()
}
