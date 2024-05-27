package config

import (
	"github.com/Pasca11/GoBlogApp/internal/api/server"
	"github.com/Pasca11/GoBlogApp/internal/pkg/logger"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
)

const envPath = ".env"

type Config struct {
	App    *App
	Server *server.Config
	Logger *logger.Config
}

type App struct {
	Name    string `yaml:"name"`
	Version string `yaml:"ver"`
}

func New() (*Config, error) {
	err := godotenv.Load(envPath)
	if err != nil {
		return nil, err
	}

	cfg := new(Config)

	err = cleanenv.ReadConfig(os.Getenv("CONFIG_PATH"), &cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
