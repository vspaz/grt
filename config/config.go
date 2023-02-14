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
	Redis    *redis.Options
	Http     *Http
	RabbitMq string
}

type Server struct {
	HostAndPort             string
	ReadTimeout             time.Duration
	WriteTimeout            time.Duration
	IdleTimeout             time.Duration
	RequestExecutionTimeout time.Duration
}

type Retries struct {
	Count    int
	Delay    float64
	OnErrors []int
}

type Timeouts struct {
	Response    float64
	Connection  float64
	HeadersRead float64
}

type Client struct {
	Host      string
	UserAgent string
	Retries   *Retries
	Timeouts  *Timeouts
}

type Http struct {
	*Client
	*Server
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
			Http: &Http{
				Client: &Client{
					Host:      "",
					UserAgent: "grt",
					Retries: &Retries{
						Count:    3,
						Delay:    0.2,
						OnErrors: []int{500, 503},
					},
					Timeouts: &Timeouts{
						Connection:  10,
						Response:    10,
						HeadersRead: 5,
					},
				},
				Server: &Server{
					HostAndPort:             ":8080",
					ReadTimeout:             10 * time.Second,
					WriteTimeout:            10 * time.Second,
					IdleTimeout:             10 * time.Second,
					RequestExecutionTimeout: 10 * time.Second,
				},
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
