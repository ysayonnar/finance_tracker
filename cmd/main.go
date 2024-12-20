package main

import (
	"financeTracker/internal/storage"
	"financeTracker/pkg/config"
	"financeTracker/pkg/logger"
	"log"
)

func main() {
	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatalln(err)
	}

	log := logger.NewLogger(cfg)

	s, err := storage.NewStorage(&cfg.DbConfig)
	if err != nil {
		log.Error("error while connecting to db", logger.CustomError(err))
	}

	_ = s
}
