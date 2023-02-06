package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/vspaz/grt/config"
	"github.com/vspaz/simplelogger/pkg/logging"
	"net/http"
	"strconv"
	"time"
)

type Router struct {
	Logger *logrus.Logger
	Conf   *config.Conf
}

func (r *Router) Get(response http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	okBody, _ := json.Marshal(map[string]string{id: "ok"})
	response.WriteHeader(http.StatusOK)
	response.Header().Set("Content-Type", "application/json")
	response.Header().Set("Content-Length", strconv.Itoa(len(okBody)))
	response.Write(okBody)
}

func RegisterHandlers(mux *chi.Mux) *chi.Mux {
	// apiV1Prefix := "/api/v1/"
	mux.Get("/ping/", Router{}.GetHealthStatus)
	mux.Handle("/metrics/", promhttp.Handler())
	logging.GetTextLogger().Logger.Info("handlers are registered: 'ok'.")
	return mux
}

func (r *Router) StartServer(mux *chi.Mux) {
	grtServer := &http.Server{
		Addr:         ":8080",
		Handler:      http.TimeoutHandler(mux, time.Second*10, "timeout occurred"),
		ReadTimeout:  10,
		WriteTimeout: 10,
		IdleTimeout:  10,
	}
	if err := grtServer.ListenAndServe(); err != nil {
		r.Logger.Fatalf("error occurred: %s")
	}
}
