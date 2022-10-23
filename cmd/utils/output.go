package utils

import (
	"fmt"
	"os"

	"github.com/guionardo/go-dev/pkg/logger"
)

var (
	OutputFileName string
)

func SetOutput(fileName string) {
	OutputFileName = fileName
	if _, err := os.Stat(OutputFileName); err == nil {
		os.Remove(OutputFileName)
	}
}

func WriteOutput(commands string) {
	if len(OutputFileName) == 0 {
		logger.Info("No output file defined")
		return
	}
	outputContent := fmt.Sprintf("#!/usr/bin/env bash\n%s", commands)
	if err := os.WriteFile(OutputFileName, []byte(outputContent), 0744); err != nil {
		logger.Fatal("Failed to write output file: %v", err)
	}
}
