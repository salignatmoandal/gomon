package core

import (
	"log"
	"os"
)

var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func LogoInfo(msg string) {
	infoLogger.Println(msg)
}

func LogoError(err error) {
	errorLogger.Println(err)
}
