package configs

import (
	"os"

	"github.com/joho/godotenv"
)

const (
	EnvVarAPIKey = "API_KEY"
)

type Configs struct {
	APIKey string
}

func Load(files ...string) (Configs, error) {
	if err := godotenv.Load(files...); err != nil {
		return Configs{}, err
	}

	configs := Configs{
		APIKey: os.Getenv(EnvVarAPIKey),
	}

	return configs, nil
}
