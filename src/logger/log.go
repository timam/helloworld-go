package logger

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

func ConfigLogger(logFile string) {
		file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		}

		log.SetOutput(file)
		log.SetFormatter(&log.JSONFormatter{
			PrettyPrint: true,
		})
		log.SetLevel(log.InfoLevel)
}