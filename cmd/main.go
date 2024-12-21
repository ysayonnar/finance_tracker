package main

import (
	"financeTracker/internal/routers"
	"financeTracker/internal/storage"
	"financeTracker/pkg/config"
	"financeTracker/pkg/logger"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
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

	r := routers.NewMainRouter(log, s)
	server := &http.Server{
		Addr:         cfg.HttpServer.Address,
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Info("Server listens", slog.String("host", cfg.HttpServer.Address))
	if err = server.ListenAndServe(); err != nil {
		log.Error("Error while starting server", logger.CustomError(err))
	}
	log.Error("Server stopped!")

}
