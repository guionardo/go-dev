package logger

import (
	"fmt"
	"os"

	"github.com/guionardo/go-dev/pkg/colors"
)

var debugMode = false

const (
	errorLabel = "[ERROR] "
	infoLabel  = ""
	debugLabel = "[DEBUG] "
	fatalLabel = "[FATAL] "
)

func SetDebugMode(mode bool) {
	debugMode = mode
}

func IsDebugMode() bool {
	return debugMode
}

func log(level string, format string, startColor string, args ...interface{}) {
	fmt.Printf(startColor+level+format+colors.Reset+"\n", args...)
}

func Debug(format string, args ...interface{}) {
	if debugMode {
		log(debugLabel, format, colors.Blue, args...)
	}
}

func Info(format string, args ...interface{}) {
	log(infoLabel, format, colors.Green, args...)
}

func Error(format string, args ...interface{}) {
	log(errorLabel, format, colors.Yellow, args...)
}

func Fatal(format string, args ...interface{}) {
	log(fatalLabel, format, colors.Red, args...)
	os.Exit(1)
}
