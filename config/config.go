package config

import (
	"github.com/go-redis/redis/v8"
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
	return SingletonConfig{}
}

func GetGlobalConfig() SingletonConfig {
	once.Do(
		func() {
			config = initConfig()
		},
	)
	return config
}
