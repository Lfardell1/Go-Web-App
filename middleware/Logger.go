package middleware

import (
	"log"
	"net/http"
	"os"
	"time"
)

// Logger is a logging helper to log important information.
type Logger struct {
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
}

// NewLogger creates a new Logger instance.
func NewLogger() *Logger {
	return &Logger{
		InfoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		ErrorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// LogInfo logs informational messages.
func (l *Logger) LogInfo(message string) {
	l.InfoLogger.Println(message)
}

// LogError logs error messages.
func (l *Logger) LogError(err error) {
	if err != nil {
		l.ErrorLogger.Println(err)
	}
}

// LogServerRequest logs information about incoming requests.
func (l *Logger) LogServerRequest(r *http.Request) {
	logMessage := time.Now().Format("2006-01-02 15:04:05") + " - " + r.Method + " " + r.URL.Path
	l.InfoLogger.Println(logMessage)
}
