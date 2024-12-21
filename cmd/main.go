package main

import (
	"financeTracker/internal/storage"
	"financeTracker/pkg/config"
	"financeTracker/pkg/logger"
	"log"
	"os"
)

func main() {
	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatalln(err)
	}

	log := logger.NewLogger(cfg)
	log.Info("Config parsed")

	s, err := storage.NewStorage(&cfg.DbConfig)
	if err != nil {
		log.Error("error while connecting to db", logger.CustomError(err))
		os.Exit(1)
	}
	log.Info("Database connected")

	_ = s
}
