package logger

import "log"

func InitLogger() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Logger initialized")
}
