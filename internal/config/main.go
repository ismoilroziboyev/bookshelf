package config

import (
	"os"

	"github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

const (
	MODE_PRODUCTION  = "PRODUCTION"
	MODE_DEVELOPMENT = "DEVELOPMENT"
)

type Config struct {
	Mode string `env:"PROJECT_MODE"`

	HttpPort string `env:"HTTP_PORT"`
	HttpHost string `env:"HTTP_HOST"`

	PSQLUri string `env:"PSQL_URI"`
}

func Load() *Config {
	var cfg Config

	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		logrus.Fatal("cannot find .env file to load environment variables")
	}

	if _, err := env.UnmarshalFromEnviron(&cfg); err != nil {
		logrus.Fatalf("error occured while parsing additional configs from environment: %s", err.Error())
	}

	return &cfg
}
