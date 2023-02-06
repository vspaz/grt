package cmd

import (
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/vspaz/simplelogger/pkg/logging"
)

type CmdArguments struct {
	LogLevel string
}

func debugCmdArgs(logger *logrus.Logger) {
	flag.VisitAll(func(flag *flag.Flag) {
		logger.Debugf("%-18s => '%v'\n", flag.Name, flag.Value)
	})
}

func GetCmdArguments(args []string) *CmdArguments {
	logLevel := flag.String("loglevel", "debug", "log level e.g. [panic | fatal | error | warning | info | debug | trace]")
	flag.Parse()
	logger := logging.GetTextLogger(*logLevel).Logger
	logger.Info("logger initialized: 'ok'.")
	flag.Usage = func() {
		logger.Printf("Help for %s:\n", args[0])
		flag.PrintDefaults()
	}
	debugCmdArgs(logger)
	logger.Info("cmd arguments are read & parsed: 'ok'.")
	return &CmdArguments{
		LogLevel: *logLevel,
	}
}
