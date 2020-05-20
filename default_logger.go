package glogger

import (
	"fmt"
	"log"
	"os"
)

var defaultLogger ILogger

const (
	defaultLogFlags = log.Llongfile | log.Ldate | log.Ltime | log.Lmicroseconds
	defaultLogLevel = LLevelDebug
)

func init() {
	defaultLogger = NewWriterLoggerWithWriter(os.Stderr)
}

func GetDefaultLogger() ILogger {
	return defaultLogger
}

func SetDefaultLogger(logger ILogger) {
	defaultLogger = logger
}

func Debug(values map[string]interface{}) {
	defaultLogger.Debug(values)
}
func Debugf(format string, a ...interface{}) {
	defaultLogger.Debugf(format, a...)
}
func Info(values map[string]interface{}) {
	defaultLogger.Info(values)
}
func Infof(format string, a ...interface{}) {
	defaultLogger.Infof(format, a...)
}
func Warn(values map[string]interface{}) {
	defaultLogger.Warn(values)
}
func Warnf(format string, a ...interface{}) {
	defaultLogger.Warnf(format, a...)
}
func Error(values map[string]interface{}) {
	defaultLogger.Error(values)
}
func Errorf(format string, a ...interface{}) {
	defaultLogger.Errorf(format, a...)
}
func Fatal(values map[string]interface{}) {
	defaultLogger.Fatal(values)
}
func Fatalf(format string, a ...interface{}) {
	defaultLogger.Fatalf(format, a...)
}

func defaultLevelLoggerPrefix(level string) string {
	return fmt.Sprintf("[%s] ", level)
}
