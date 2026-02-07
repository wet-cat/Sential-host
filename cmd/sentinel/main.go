package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"sentinel/internal/logger"
	"sentinel/internal/nfqueue"
)

func main() {
	safe := flag.Bool("safe", true, "enable logging-only mode; do not drop packets")
	flag.Parse()

	logger.Info("starting sentinel-host (safe mode=" + boolToStr(*safe) + ")")

	if err := nfqueue.Run(*safe); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	logger.Info("shutdown requested, exiting")
}

func boolToStr(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
