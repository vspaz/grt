package handlers

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/vspaz/grt/config"
	"github.com/vspaz/simplelogger/pkg/logging"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
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

func (r *Router) handleShutDownGracefully(server *http.Server) {
	signals := []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, signals...)
	r.Logger.Infof("'%s' signal received, stopping server...", <-signalChannel)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		r.Logger.Errorf("error shutting down server: %s", err)
	}
}

func (r *Router) StartServer(mux *chi.Mux) {
	server := &http.Server{
		Addr:         r.Conf.HttpServer.HostAndPort,
		Handler:      http.TimeoutHandler(mux, r.Conf.HttpServer.RequestExecutionTimeout, "timeout occurred"),
		ReadTimeout:  r.Conf.HttpServer.ReadTimeout,
		WriteTimeout: r.Conf.HttpServer.WriteTimeout,
		IdleTimeout:  r.Conf.HttpServer.IdleTimeout,
	}
	go r.handleShutDownGracefully(server)
	pid := os.Getpid()
	r.Logger.Infof("starting server pid='%d' at port %s.", pid, r.Conf.HttpServer.HostAndPort)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		r.Logger.Fatalf("error occurred: %s", err)
	}
	r.Logger.Info("server stopped.")
}
