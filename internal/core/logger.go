package core

import (
	"log"
	"os"
)

type Logger struct {
	Level   string
	Message string
	Error   error
}

var logCh = make(chan Logger, 100)

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func init() {
	go func() {
		for entry := range logCh {
			if entry.Level == "info" {
				infoLogger.Println(entry.Message)
			} else if entry.Level == "error" && entry.Error != nil {
				errorLogger.Println(entry.Error)
			}
		}
	}()
}

func LogInfo(msg string) {
	logCh <- Logger{Level: "info", Message: msg}
}

func LogError(err error) {
	logCh <- Logger{Level: "error", Error: err}
}
