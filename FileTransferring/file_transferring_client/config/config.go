package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App `yaml:"app"`
		Log `yaml:"logger"`
		Network
		Data
	}

	App struct {
		Name    string `yaml:"name"    env:"APP_NAME" env-default:"File transferring client"`
		Version string `yaml:"version" env:"APP_VERSION" env-default:"1.0.0"`
	}

	Network struct {
		IpAddress  string
		Port       string
		DataLength int `yaml:"data_length" env:"DATA_LENGTH" env-default:"1024"`
	}

	Data struct {
		FilePath string
	}

	Log struct {
		Level string `yaml:"log_level"   env:"LOG_LEVEL" env-default:"info"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
