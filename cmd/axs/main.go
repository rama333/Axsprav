package main

import (
	"Axsprav/internal/config"
	"Axsprav/internal/resources"
	"Axsprav/internal/restapi"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initZapLog() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, _ := config.Build()
	return logger
}

func main()  {
	logger := initZapLog()
	defer logger.Sync()

	config.Config.LOGGER = logger.Sugar()
	slogger := config.Config.LOGGER

	slogger.Info("Starting the application...")
	slogger.Info("Reading configuration and initializing resources...")

	if err := config.LoadConfig("/smsService"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	_, err := resources.New(slogger)
	if err != nil {
		slogger.Fatalw("Can't initialize resources.", "err", err)
	}

	slogger.Info("Configuring the application units...")

	//diag := diagnostics.New(slogger, config.Config.DIAGPORT, rsc.Healthz)
	//diag.Start(slogger)
	//slogger.Info("The application is ready to serve requests.")

	rapi := restapi.New(slogger)
	rapi.Start(config.Config.RESTAPIPort)
}
