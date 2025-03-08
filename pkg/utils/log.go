package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// LogToFile creates a log file with today's date in the `storage/logs` directory
// if it doesn't exist, and logs the provided parameters.
func LogToFile(params ...interface{}) error {
	// Get the current date in YYYY-MM-DD format
	today := time.Now().Format("2006-01-02")

	// Define the log file path
	logDir := "storage/logs"
	logFilePath := filepath.Join(logDir, today+".log")

	// Create the directory if it doesn't exist
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.MkdirAll(logDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create log directory: %w", err)
		}
	}

	// Open the log file (create it if it doesn't exist, append to it if it does)
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	defer logFile.Close()

	// Create a logger to write to the file
	logger := log.New(logFile, "", log.LstdFlags)

	// Log the parameters as a formatted string
	logger.Printf(fmt.Sprint(params...))

	return nil
}
