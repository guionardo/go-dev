package utils

import (
	"fmt"
	"log"
	"os"
)

var (
	OutputFileName string
)

func SetOutput(fileName string) {
	OutputFileName = fileName
}

func WriteOutput(commands string) {
	outputContent := fmt.Sprintf("#!/bin/bash\n%s", commands)
	if err := os.WriteFile(OutputFileName, []byte(outputContent), 0744); err != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}
}
