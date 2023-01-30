package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

func ConfigureMiddleware(logger *logrus.Logger) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	return router
}
