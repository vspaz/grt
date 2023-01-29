package config

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"sync"
)

var (
	config SingletonConfig
	once   sync.Once
)

type SingletonConfig struct {
	Config *Conf
}

type Conf struct {
	Redis *redis.Options
}

func initConfig() SingletonConfig {
	return SingletonConfig{
		Config: &Conf{
			Redis: &redis.Options{
				Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
				Password: os.Getenv("REDIS_PASSWORD"),
			},
		},
	}
}

func GetConfig() SingletonConfig {
	once.Do(
		func() {
			config = initConfig()
		},
	)
	return config
}
