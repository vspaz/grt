package handlers

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/vspaz/grt/config"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type Router struct {
	Logger      *logrus.Logger
	Conf        *config.Conf
	RedisClient *redis.Client
	redisCtx    context.Context
	mux         *chi.Mux
}

func NewRouter(conf *config.Conf, logger *logrus.Logger) *Router {
	return &Router{
		Logger: logger,
		Conf:   conf,
		mux:    chi.NewRouter(),
	}
}

func (r *Router) SetRedisClient(client *redis.Client) {
	r.RedisClient = client
	r.redisCtx = context.Background()
	r.Logger.Info("redis client initialized: 'ok'")
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

func (r *Router) RegisterHandlers() {
	// apiV1Prefix := "/api/v1/"
	r.mux.Get("/ping/", Router{}.GetHealthStatus)
	r.mux.Handle("/metrics/", promhttp.Handler())
	r.Logger.Info("handlers are registered: 'ok'.")
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

func (r *Router) StartServer() {
	server := &http.Server{
		Addr:         r.Conf.Http.Server.HostAndPort,
		Handler:      http.TimeoutHandler(r.mux, r.Conf.Http.Server.RequestExecutionTimeout, "timeout occurred"),
		ReadTimeout:  r.Conf.Http.Server.ReadTimeout,
		WriteTimeout: r.Conf.Http.Server.WriteTimeout,
		IdleTimeout:  r.Conf.Http.Server.IdleTimeout,
	}
	go r.handleShutDownGracefully(server)
	r.Logger.Infof("starting server pid='%d' at port '%s'.", os.Getpid(), r.Conf.Http.Server.HostAndPort)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		r.Logger.Fatalf("error occurred: %s", err)
	}
	r.Logger.Info("server stopped.")
}
