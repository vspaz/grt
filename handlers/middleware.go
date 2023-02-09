package handlers

import (
	"github.com/chi-middleware/logrus-logger"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func (r *Router) ConfigureMiddleware() {
	r.mux.Use(middleware.Recoverer)
	r.mux.Use(middleware.RealIP)
	r.mux.Use(middleware.RequestID)
	r.mux.Use(logger.Logger("http-loggger", r.Logger))
	r.mux.Use(render.SetContentType(render.ContentTypeJSON))
	r.Logger.Info("middleware is configured: 'ok'.")
}
