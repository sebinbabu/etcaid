package main

import (
	"os"

	"github.com/sebinbabu/etcaid"
	"github.com/sebinbabu/etcaid/simplelogger"
)

// buildController builds & returns an etcaid controller for use with the cli
// by creating its dependencies first and using them to build one.
func buildController() *etcaid.Controller {
	logger := simplelogger.New()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		logger.Error("failed to find user home dir")
		os.Exit(1)
	}

	confDir, err := os.UserConfigDir()
	if err != nil {
		logger.Error("failed to find user config dir")
		os.Exit(1)
	}

	controller := etcaid.NewController(
		homeDir,
		confDir,
		logger,
	)

	return controller
}
