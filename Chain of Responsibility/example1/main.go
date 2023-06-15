package main

import "fmt"

type Logger interface {
	SetNext(Logger)
	LogMessage(int, string)
}

type AbstractLogger struct {
	Level      int
	NextLogger Logger
	WriteFunc  func(string)
}

func (l *AbstractLogger) SetNext(nextLogger Logger) {
	l.NextLogger = nextLogger
}

func (l *AbstractLogger) LogMessage(level int, message string) {
	if l.Level <= level {
		l.WriteFunc(message)
	}
	if l.NextLogger != nil {
		l.NextLogger.LogMessage(level, message)
	}
}

func (l *AbstractLogger) Write(message string) {
	panic("write method must be implemented by concrete loggers")
}

type ErrorLogger struct {
	AbstractLogger
}

func (l *ErrorLogger) Write(message string) {
	fmt.Printf("Error: %s\n", message)
}

func NewErrorLogger() *ErrorLogger {
	l := &ErrorLogger{AbstractLogger{Level: 1}}
	l.WriteFunc = l.Write
	return l
}

type DebugLogger struct {
	AbstractLogger
}

func (l *DebugLogger) Write(message string) {
	fmt.Printf("Debug: %s\n", message)
}

func NewDebugLogger() *DebugLogger {
	l := &DebugLogger{AbstractLogger{Level: 2}}
	l.WriteFunc = l.Write
	return l
}

type InfoLogger struct {
	AbstractLogger
}

func (l *InfoLogger) Write(message string) {
	fmt.Printf("Info: %s\n", message)
}

func NewInfoLogger() *InfoLogger {
	l := &InfoLogger{AbstractLogger{Level: 3}}
	l.WriteFunc = l.Write
	return l
}

func main() {
	errorLogger := NewErrorLogger()
	debugLogger := NewDebugLogger()
	infoLogger := NewInfoLogger()

	errorLogger.SetNext(debugLogger)
	debugLogger.SetNext(infoLogger)

	errorLogger.LogMessage(1, "This is an error message")
	errorLogger.LogMessage(2, "This is a debug message")
	errorLogger.LogMessage(3, "This is an info message")
}
