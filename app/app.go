package app

import (
	"github.com/go-redis/redis/v8"
	"github.com/vspaz/grt/cmd"
	"github.com/vspaz/grt/config"
	"github.com/vspaz/grt/handlers"
	"github.com/vspaz/simplelogger/pkg/logging"
	"os"
)

func Run(binaryName string) {
	args := cmd.GetCmdArguments(os.Args)
	conf := config.GetConfig().Config
	logger := logging.GetTextLogger(args.LogLevel).Logger
	logger.Infof("grt server build='%s'", binaryName)
	router := handlers.NewRouter(logger, conf)
	router.SetRedisClient(redis.NewClient(conf.Redis))
	mux := handlers.ConfigureMiddleware(logger)
	mux = handlers.RegisterHandlers(mux)
	router.StartServer(mux)
}
