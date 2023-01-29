package app

import (
	"github.com/vspaz/grt/cmd"
	"github.com/vspaz/grt/config"
	"github.com/vspaz/simplelogger/pkg/logging"
	"os"
)

func Run() {
	args := cmd.GetCmdArguments(os.Args)
	globalConfig := config.GetConfig().Config
	logger := logging.GetTextLogger(args.LogLevel).Logger
	logger.Info("app started")
	logger.Info(globalConfig)
}
