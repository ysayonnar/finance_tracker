package userHandlers

import (
	"financeTracker/internal/storage"
	"fmt"
	"log/slog"
	"net/http"
)

func Registration(log *slog.Logger, s *storage.Storage) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const op = `userHandlers.Registration`

		fmt.Fprintln(w, op)
	})
}

func Login(log *slog.Logger, s *storage.Storage) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const op = `userHandlers.Login`

		fmt.Fprintln(w, op)
	})
}
