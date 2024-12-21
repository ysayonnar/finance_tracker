package routers

import (
	userHandlers "financeTracker/internal/handlers/user"
	"financeTracker/internal/storage"
	"log/slog"

	"github.com/gorilla/mux"
)

func RegisterUserRouter(r *mux.Router, log *slog.Logger, s *storage.Storage) {
	userRouter := r.PathPrefix("/user").Subrouter()

	userRouter.Handle("/registration", userHandlers.Registration(log, s)).Methods("POST")
	userRouter.Handle("/login", userHandlers.Login(log, s)).Methods("POST")
}
