package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	defaultServerPort    = "8080"
	defaultServerHost    = "http://localhost:8080"
	defaultServerTimeout = 60 * time.Second

	defaultTokenKey     = "IP03O5Ekg91g5jw=="
	defaultTokenExpires = 3600 * time.Second
)

type (
	Configs struct {
		SERVER ServerConfig
		TOKEN  TokenConfig
	}

	ServerConfig struct {
		Port    string        `required:"true"`
		Host    string        `required:"true"`
		Timeout time.Duration `required:"true"`
	}

	TokenConfig struct {
		Key     string
		Expires time.Duration
	}

	ClientConfig struct {
		URL      string `required:"true"`
		Login    string
		Password string
	}
)

// New populates Configs struct with values from config file
// located at filepath and environment variables.
func New() (cfg Configs, err error) {
	root, err := os.Getwd()
	if err != nil {
		return
	}
	godotenv.Load(filepath.Join(root, ".env"))

	cfg.SERVER = ServerConfig{
		Port:    defaultServerPort,
		Host:    defaultServerHost,
		Timeout: defaultServerTimeout,
	}

	cfg.TOKEN = TokenConfig{
		Key:     defaultTokenKey,
		Expires: defaultTokenExpires,
	}

	if err = envconfig.Process("SERVER", &cfg.SERVER); err != nil {
		return
	}

	return
}
