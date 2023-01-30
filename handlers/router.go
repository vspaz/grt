package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type Router struct {
	Logger *logrus.Logger
}

func RegisterHandlers(mux *chi.Mux, handlers *Router) *chi.Mux {
	// apiV1Prefix := "/api/v1/"
	mux.Get("/ping/", Router{}.GetHealthStatus)
	return mux
}
