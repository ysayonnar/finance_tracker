package routers

import (
	mw "financeTracker/internal/middlewares"
	"financeTracker/internal/storage"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

func NewMainRouter(log *slog.Logger, s *storage.Storage) *mux.Router {
	r := mux.NewRouter()
	r.Use(func(next http.Handler) http.Handler {
		return mw.RequestStats(next, log)
	})

	RegisterUserRouter(r, log, s)
	//TODO: точно так же регестрировать здесь остальные роутеры

	return r
}
