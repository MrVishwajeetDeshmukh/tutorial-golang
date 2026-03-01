package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.InfoLevel)

	log.Debug("This message will not be printed.")
	log.Info("This is an info message.")
	log.Warn("This is a warning message.")
	log.Error("This is an error message.")
}
