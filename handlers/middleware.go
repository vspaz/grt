package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

func ConfigureMiddleware(log *logrus.Logger) *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.RequestID)
	mux.Use(middleware.Heartbeat("/ping"))
	mux.Use(render.SetContentType(render.ContentTypeJSON))
	return mux
}
