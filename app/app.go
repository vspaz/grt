package app

import (
	"github.com/vspaz/simplelogger/pkg/logging"
	"gtf/cmd"
	"os"
)

func Run() {
	args := cmd.GetCmdArguments(os.Args)
	logger := logging.GetTextLogger(args.LogLevel).Logger
	logger.Info("app started")
}
