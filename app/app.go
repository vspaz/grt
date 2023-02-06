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
	router := handlers.Router{
		Logger: logger,
		Conf:   globalConfig,
	}
	logger.Infof("server build, ver='%s'", binaryName)
	mux := handlers.ConfigureMiddleware(logger)
	mux = handlers.RegisterHandlers(mux)
	router.StartServer(mux)
}
