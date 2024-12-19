package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type ConfigValidationError struct {
	Message string
}

func (e ConfigValidationError) Error() string {
	return fmt.Sprintf("Error while validating config: %s", e.Message)
}

type HttpServer struct {
	Address string `json:"address"`
}

type DatabaseConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"db_name"`
	SslMode  string `json:"ssl_mode"`
}

type Config struct {
	Env        string         `json:"env"`
	HttpServer HttpServer     `json:"http_server"`
	DbConfig   DatabaseConfig `json:"db"`
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

	//TODO: валидация пустых полей через go-validator

	return &cfg, nil
}
