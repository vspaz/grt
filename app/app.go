package app

import (
	"github.com/vspaz/simplelogger/pkg/logging"
	"gtf/cmd"
	"gtf/config"
	"os"
)

func Run() {
	args := cmd.GetCmdArguments(os.Args)
	globalConfig := config.GetConfig().Config
	logger := logging.GetTextLogger(args.LogLevel).Logger
	logger.Info("app started")
}
