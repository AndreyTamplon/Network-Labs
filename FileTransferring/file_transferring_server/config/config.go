package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App `yaml:"app"`
		Log `yaml:"logger"`
		Network `yaml:"network"`

	}

	App struct {
		Name    string `yaml:"name"    env:"APP_NAME" env-default:"File transferring server"`
		Version string `yaml:"version" env:"APP_VERSION" env-default:"1.0.0"`
	}

	Network struct {
		Port      string `yaml:"port"      env:"APP_PORT" env-default:"8084"`
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
