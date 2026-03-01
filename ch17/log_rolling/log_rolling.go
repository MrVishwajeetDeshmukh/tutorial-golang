package main

import (
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    5, // Megabytes
		MaxBackups: 3,
		MaxAge:     28,   // Days
		Compress:   true, // Whether to compress
	})

	for i := 0; i < 1000000; i++ {
		log.Printf("Log message %d", i)
	}
}
