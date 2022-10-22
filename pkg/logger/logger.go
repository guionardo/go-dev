package logger

import "fmt"

var debugMode = false

func SetDebugMode(mode bool) {
	debugMode = mode
}

func log(level string, format string, args ...interface{}) {
	fmt.Printf("["+level+"] "+format+"\n", args...)
}

func Debug(format string, args ...interface{}) {
	if debugMode {
		log("DEBUG", format, args...)
	}
}

func Info(format string, args ...interface{}) {
	log("INFO ", format, args...)
}
