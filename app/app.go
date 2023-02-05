package app

import (
	"github.com/vspaz/grt/cmd"
	"github.com/vspaz/grt/config"
	"github.com/vspaz/grt/handlers"
	"github.com/vspaz/simplelogger/pkg/logging"
	"os"
)

func Run(binaryName string) {
	args := cmd.GetCmdArguments(os.Args)
	globalConfig := config.GetConfig().Config
	logger := logging.GetTextLogger(args.LogLevel).Logger
	pid := os.Getpid()
	logger.Infof("starting server pid='%d'", pid)
	logger.Infof("server build, ver='%s'", binaryName)
	logger.Info(globalConfig)
	mux := handlers.ConfigureMiddleware(logger)
	mux = handlers.RegisterHandlers(mux)
	router := handlers.Router{}
	router.StartServer(mux)
}
