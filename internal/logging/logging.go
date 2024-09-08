package logging

import (
	"log"
	"os"
)

var (
	Logger *log.Logger
)

// InitializeLogger sets up the logger to write to a specified file.
func InitializeLogger(logFilePath string) {
	// Open or create the log file
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}

	// Set the global logger to use this file
	Logger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}
