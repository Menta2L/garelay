package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/mhale/smtpd"
	"go.uber.org/zap"
)

var (
	cfg config
	logger *zap.Logger
)

func main() {
	cfg = config{}
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		fmt.Println("Error cannot initialize logger")
	}
	defer logger.Sync() // flushes buffer, if any
	if err := env.Parse(&cfg); err != nil {
		logger.Error("Failed to load configuration",
			zap.String("error", err.Error()),
		)
	}
	address := fmt.Sprintf("%s:%d",cfg.Address,cfg.Port)
	err= smtpd.ListenAndServe(address, mailHandler, "GA-Proxy", "")
	if err != nil {
		logger.Error("Failed to bind address ",
			zap.String("err", err.Error()),
		)
	}
}