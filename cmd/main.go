package main

import (
	"financeTracker/pkg/config"
	"financeTracker/pkg/logger"
	"log"
)

func main() {
	//config initialization
	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatalln(err)
	}

	//logger config
	log := logger.NewLogger(cfg)
	_ = log

	customError := config.ConfigValidationError{Message: "ты даун"}
	log.Error("error", logger.CustomError(customError))
}
