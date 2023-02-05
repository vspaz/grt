package app

import (
	"github.com/vspaz/grt/cmd"
	"github.com/vspaz/grt/config"
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
	logger.Info("app started")
	logger.Info(globalConfig)
}
