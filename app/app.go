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
	logger.Infof("grt server build, ver='%s'", binaryName)
	router := handlers.Router{
		Logger: logger,
		Conf:   globalConfig,
	}
	mux := handlers.ConfigureMiddleware(logger)
	mux = handlers.RegisterHandlers(mux)
	router.StartServer(mux)
}
