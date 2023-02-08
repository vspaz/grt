package config

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"sync"
	"time"
)

var (
	config *SingletonConfig
	once   sync.Once
)

type SingletonConfig struct {
	Config *Conf
}

type Conf struct {
	Redis      *redis.Options
	HttpServer *HttpServer
	RabbitMq   string
}

type Server struct {
	HostAndPort             string
	ReadTimeout             time.Duration
	WriteTimeout            time.Duration
	IdleTimeout             time.Duration
	RequestExecutionTimeout time.Duration
}

type Client struct {
}

type Http struct {
	Server
	Client
}

type HttpServer struct {
	HostAndPort             string
	ReadTimeout             time.Duration
	WriteTimeout            time.Duration
	IdleTimeout             time.Duration
	RequestExecutionTimeout time.Duration
}

func initConfig() *SingletonConfig {
	return &SingletonConfig{
		Config: &Conf{
			Redis: &redis.Options{
				Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
				Password: os.Getenv("REDIS_PASSWORD"),
			},
			HttpServer: &HttpServer{
				HostAndPort:             ":8080",
				ReadTimeout:             10 * time.Second,
				WriteTimeout:            10 * time.Second,
				IdleTimeout:             10 * time.Second,
				RequestExecutionTimeout: 10 * time.Second,
			},
			RabbitMq: "",
		},
	}
}

func GetConfig() *SingletonConfig {
	once.Do(
		func() {
			config = initConfig()
		},
	)
	return config
}
