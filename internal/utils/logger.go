/*
 **********************************************************************
 * -------------------------------------------------------------------
 * Project Name : Abdal Better PassList
 * File Name    : logger.go
 * Author       : Ebrahim Shafiei (EbraSha)
 * Email        : Prof.Shafiei@Gmail.com
 * Created On   : 2025-09-16 15:40:00
 * Description  : Error logging system for the application
 * -------------------------------------------------------------------
 *
 * "Coding is an engaging and beloved hobby for me. I passionately and insatiably pursue knowledge in cybersecurity and programming."
 * – Ebrahim Shafiei
 *
 **********************************************************************
 */

package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

var (
	errorLogger *log.Logger
	logFile     *os.File
)

// InitLogger initializes the error logging system
func InitLogger() error {
	// Create logs directory if it doesn't exist
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return fmt.Errorf("failed to create logs directory: %v", err)
	}

	// Create log file with timestamp
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	logFileName := fmt.Sprintf("abdal_passlist_errors_%s.log", timestamp)
	logFilePath := filepath.Join(logDir, logFileName)

	// Open log file
	var err error
	logFile, err = os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}

	// Create logger
	errorLogger = log.New(logFile, "", 0)

	// Log initialization (as info, not error)
	LogInfo("Logger initialized successfully")
	return nil
}

// LogInfo logs an info message to the log file
func LogInfo(message string) {
	if errorLogger == nil {
		// If logger is not initialized, just print to console
		fmt.Printf("INFO: %s\n", message)
		return
	}

	// Get caller information
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}

	// Extract just the filename from the full path
	fileName := filepath.Base(file)

	// Format timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// Create log entry
	logEntry := fmt.Sprintf("[%s] INFO in %s:%d - %s", timestamp, fileName, line, message)

	// Write to log file
	errorLogger.Println(logEntry)
}

// LogError logs an error to the log file
func LogError(message string, err error) {
	if errorLogger == nil {
		// If logger is not initialized, try to initialize it
		if initErr := InitLogger(); initErr != nil {
			// If we can't initialize logger, just print to console
			fmt.Printf("ERROR: %s - %v (Logger init failed: %v)\n", message, err, initErr)
			return
		}
	}

	// Get caller information
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}

	// Extract just the filename from the full path
	fileName := filepath.Base(file)

	// Format timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// Create log entry
	var logEntry string
	if err != nil {
		logEntry = fmt.Sprintf("[%s] ERROR in %s:%d - %s: %v", timestamp, fileName, line, message, err)
	} else {
		logEntry = fmt.Sprintf("[%s] ERROR in %s:%d - %s", timestamp, fileName, line, message)
	}

	// Write to log file
	errorLogger.Println(logEntry)

	// Also print to console for immediate feedback
	fmt.Printf("ERROR: %s\n", logEntry)
}

// LogPanic logs a panic and recovers from it
func LogPanic(message string) {
	if r := recover(); r != nil {
		LogError(fmt.Sprintf("PANIC: %s", message), fmt.Errorf("panic: %v", r))
		panic(r) // Re-panic to maintain normal panic behavior
	}
}

// CloseLogger closes the log file
func CloseLogger() {
	if logFile != nil {
		LogInfo("Logger closing")
		logFile.Close()
	}
}

// GetLogFilePath returns the current log file path
func GetLogFilePath() string {
	if logFile != nil {
		return logFile.Name()
	}
	return "logs/abdal_passlist_errors.log"
}
