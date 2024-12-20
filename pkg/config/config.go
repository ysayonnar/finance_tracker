package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-playground/validator/v10"
)

type ConfigValidationError struct {
	Message string
}

func (e ConfigValidationError) Error() string {
	return fmt.Sprintf("Error while validating config: %s", e.Message)
}

type HttpServer struct {
	Address string `json:"address" validate:"required"`
}

type DatabaseConfig struct {
	User     string `json:"user" validate:"required"`
	Password string `json:"password" validate:"required"`
	DbName   string `json:"db_name" validate:"required"`
	SslMode  string `json:"ssl_mode" validate:"required"`
}

type Config struct {
	Env        string         `json:"env" validate:"required"`
	HttpServer HttpServer     `json:"http_server" validate:"required"`
	DbConfig   DatabaseConfig `json:"db" validate:"required"`
}

func ParseConfig() (*Config, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	cfgPath := filepath.Join(cwd, "..", "config", "config.json")

	cfgBuffer, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = json.Unmarshal(cfgBuffer, &cfg)
	if err != nil {
		return nil, err
	}

	//config validation stage
	if cfg.Env != "local" && cfg.Env != "prod" {
		return nil, ConfigValidationError{Message: "Invalid environment variable"}
	}

	validate := validator.New()
	err = validate.Struct(cfg)
	if err != nil {
		return nil, ConfigValidationError{Message: err.Error()}
	}

	return &cfg, nil
}
