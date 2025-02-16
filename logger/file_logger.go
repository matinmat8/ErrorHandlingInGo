package logger

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

var (
	logger      zerolog.Logger
	currentDate string
	mu          sync.Mutex // Mutex to protect shared resources
)

func init() {
	setupLogger()
}

func setupLogger() {
	mu.Lock()
	defer mu.Unlock()

	// Generate the filename based on the current date
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	log_file := os.Getenv("LOG_FILE")
	currentDate = time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("logs/%v-%s.log", log_file, currentDate)

	// Set up lumberjack logger for log rotation with compression
	logFile := &lumberjack.Logger{
		Filename:   filename, // Use filename with date
		MaxSize:    200,      // Maximum size in megabytes before rotation
		MaxBackups: 3,        // Maximum number of old log files to retain
		Compress:   true,     // Compress rotated log files
		LocalTime:  true,     // Use local time for file timestamps
	}

	logger = zerolog.New(logFile).With().Timestamp().Logger().Output(zerolog.ConsoleWriter{
		Out:        logFile,
		NoColor:    true,
		TimeFormat: "15:04:05",
		FormatLevel: func(i interface{}) string {
			if ll, ok := i.(string); ok {
				return strings.ToUpper(ll)
			}
			return ""
		},
	})
	fmt.Println("Logger reinitialized with filename:", filename)
}

// LogErrorWithDepth Log Error with the file path
func LogErrorWithDepth(callerDepth int, err error, msg ...string) {
	_, file, line, ok := runtime.Caller(callerDepth)
	if !ok {
		file = "unknown"
		line = 0
	}

	var message string
	if len(msg) > 0 {
		message = msg[0]
	} else {
		message = "An Error Occurred"
	}

	customMsg := fmt.Sprintf("File: %s, Line: %d, ErrorMessage: \"%s\", ErrorDetails: \"%v\"", file, line, message, err)
	logger.Error().Msg(customMsg)
}

// LogError logs an error directly at the call location before the panic
func LogError(err error, msg ...string) {
	// Retrieve caller information based on the current depth
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}

	var message string
	if len(msg) > 0 {
		message = msg[0]
	} else {
		message = "An Error Occurred"
	}

	customMsg := fmt.Sprintf("File: %s, Line: %d, ErrorMessage: \"%s\", ErrorDetails: \"%v\"", file, line, message, err)
	logger.Error().Msg(customMsg)
}
