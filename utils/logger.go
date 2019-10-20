package utils

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "LOG: ", log.Lshortfile)

func GetLogger() *log.Logger {
	return logger
}
