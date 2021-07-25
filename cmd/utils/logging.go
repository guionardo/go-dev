package utils

import (
	"io"
	"log"
	"log/syslog"
	"os"
)

var (
	sysLogger io.Writer
)

func SetupLogging() {
	var err error

	sysLogger, err = syslog.New(syslog.LOG_INFO,"go-dev")
	if err == nil {
		log.SetOutput(io.MultiWriter(sysLogger, os.Stdout))
	} else {
		log.Printf("Failed to create syslog %v\n", err)
	}
}
