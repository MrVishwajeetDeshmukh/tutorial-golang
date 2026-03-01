package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("This is a log message in JSON format.")
}
