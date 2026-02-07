package main

import (
	"os"
	"os/signal"
	"syscall"

	"sentinel/internal/logger"
	"sentinel/internal/nfqueue"
)

func main() {
	logger.Info("starting sentinel-host")

	if err := nfqueue.Run(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	logger.Info("shutdown requested, exiting")
}
