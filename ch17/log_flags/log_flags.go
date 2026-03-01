package main

import (
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Log settings have been modified.")
}
