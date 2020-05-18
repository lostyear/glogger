package glogger

import "os"

type ILogger interface {
	Debug(map[string]interface{})
	Debugf(string, ...interface{})
	Info(map[string]interface{})
	Infof(string, ...interface{})
	Warn(map[string]interface{})
	Warnf(string, ...interface{})
	Error(map[string]interface{})
	Errorf(string, ...interface{})
	Fatal(map[string]interface{})
	Fatalf(string, ...interface{})
}

type Logger struct {
	ILogger

	config *LoggerConfig

	debugLogger iLevelLogger
	infoLogger  iLevelLogger
	warnLogger  iLevelLogger
	errorLogger iLevelLogger
	fatalLogger iLevelLogger
}

func NewLoggerWithConfigFile(path string) *Logger {
	config := loadConfigFile(path)
	return NewLoggerWithConfig(*config)
}

func NewLoggerWithConfig(cfg LoggerConfig) *Logger {
	logger := Logger{
		config: &cfg,
	}
	return &logger
}

func (l *Logger) GetConfig() LoggerConfig {
	return *l.config
}

func (l *Logger) Debug(values map[string]interface{}) {
	l.debugLogger.Print(values)
}

func (l *Logger) Debugf(format string, a ...interface{}) {
	l.debugLogger.Printf(format, a...)
}

func (l *Logger) Info(values map[string]interface{}) {
	l.infoLogger.Print(values)
}

func (l *Logger) Infof(format string, a ...interface{}) {
	l.infoLogger.Printf(format, a...)
}

func (l *Logger) Warn(values map[string]interface{}) {
	l.warnLogger.Print(values)
}

func (l *Logger) Warnf(format string, a ...interface{}) {
	l.warnLogger.Printf(format, a...)
}

func (l *Logger) Error(values map[string]interface{}) {
	l.errorLogger.Print(values)
}

func (l *Logger) Errorf(format string, a ...interface{}) {
	l.errorLogger.Printf(format, a...)
}

func (l *Logger) Fatal(values map[string]interface{}) {
	l.fatalLogger.Print(values)
	os.Exit(1)
}

func (l *Logger) Fatalf(format string, a ...interface{}) {
	l.fatalLogger.Printf(format, a...)
	os.Exit(1)
}

type iLevelLogger interface {
	Print(map[string]interface{})
	Printf(string, ...interface{})
}
