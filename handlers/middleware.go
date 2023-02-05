package handlers

import (
	"github.com/chi-middleware/logrus-logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

func ConfigureMiddleware(loggger *logrus.Logger) *chi.Mux {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.RequestID)
	mux.Use(logger.Logger("http-loggger", loggger))
	mux.Use(middleware.Heartbeat("/ping"))
	mux.Use(render.SetContentType(render.ContentTypeJSON))
	loggger.Info("middleware is configured: 'ok'.")
	return mux
}
