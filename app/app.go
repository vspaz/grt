package app

import (
	"github.com/go-redis/redis/v8"
	"github.com/vspaz/goat/pkg/ghttp"
	"github.com/vspaz/grt/cmd"
	"github.com/vspaz/grt/config"
	"github.com/vspaz/grt/handlers"
	"github.com/vspaz/rmqclient/pkg/rmq"
	"github.com/vspaz/simplelogger/pkg/logging"
	"os"
)

func Run(binaryName string) {
	args := cmd.GetCmdArguments(os.Args)
	conf := config.GetConfig().Config
	logger := logging.GetTextLogger(args.LogLevel).Logger
	logger.Infof("grt server build='%s'", binaryName)
	router := handlers.NewRouter(conf, logger)
	httpClient := ghttp.NewClientBuilder().
		WithHost(conf.Http.Client.Host).
		WithUserAgent(conf.Http.Client.UserAgent).
		WithConnectionTimeout(conf.Http.Client.Timeouts.Connection).
		WithResponseTimeout(conf.Http.Client.Timeouts.Response).
		WithRetry(conf.Http.Client.Retries.Count, conf.Http.Client.Retries.OnErrors).
		WithDelay(conf.Http.Client.Retries.Delay).
		Build()
	router.SetHttpClient(httpClient)
	router.SetRedisClient(redis.NewClient(conf.Redis))
	rabbitMqConnection := rmq.NewConnection(conf.RabbitMq, logger)
	router.SetRabbitMqConnection(rabbitMqConnection)
	router.ConfigureMiddleware()
	router.RegisterHandlers()
	router.StartServer()
}
