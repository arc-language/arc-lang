package compiler

import (
	"fmt"
	"sync"
)

// LogLevel controls internal verbosity
type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelError
)

var CurrentLogLevel = LogLevelInfo

// Logger handles internal compiler messaging
type Logger struct {
	mu     sync.Mutex
	prefix string
}

func NewLogger(prefix string) *Logger {
	return &Logger{prefix: prefix}
}

func (l *Logger) Debug(format string, args ...interface{}) {
	if CurrentLogLevel <= LogLevelDebug {
		l.log("DEBUG", format, args...)
	}
}

func (l *Logger) Info(format string, args ...interface{}) {
	if CurrentLogLevel <= LogLevelInfo {
		l.log("INFO", format, args...)
	}
}

func (l *Logger) Error(format string, args ...interface{}) {
	if CurrentLogLevel <= LogLevelError {
		l.log("ERROR", format, args...)
	}
}

func (l *Logger) log(level, format string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	
	msg := fmt.Sprintf(format, args...)
	// Print to stdout/stderr with prefix
	fmt.Printf("%s [%s] %s\n", l.prefix, level, msg)
}