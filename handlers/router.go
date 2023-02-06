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
	_, err := response.Write(okBody)
	if err != nil {
		r.Logger.Errorf("error occurred: %s", err.Error())
	}
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
		Addr:         r.Conf.HttpServer.HostAndPort,
		Handler:      http.TimeoutHandler(mux, r.Conf.HttpServer.RequestExecutionTimeout, "timeout occurred"),
		ReadTimeout:  r.Conf.HttpServer.ReadTimeout,
		WriteTimeout: r.Conf.HttpServer.WriteTimeout,
		IdleTimeout:  r.Conf.HttpServer.IdleTimeout,
	}
	if err := grtServer.ListenAndServe(); err != nil {
		r.Logger.Fatalf("error occurred: %s")
	}
}
