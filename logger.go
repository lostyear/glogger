package glogger

import "os"

type ILogger interface {
	Debug(values map[string]interface{})
	Debugf(format string, a ...interface{})
	Info(values map[string]interface{})
	Infof(format string, a ...interface{})
	Warn(values map[string]interface{})
	Warnf(format string, a ...interface{})
	Error(values map[string]interface{})
	Errorf(format string, a ...interface{})
	Fatal(values map[string]interface{})
	Fatalf(format string, a ...interface{})
}

type Logger struct {
	ILogger

	config *FileLoggerConfig

	debugLogger levelLogger
	infoLogger  levelLogger
	warnLogger  levelLogger
	errorLogger levelLogger
	fatalLogger levelLogger
}

func NewFileLoggerWithConfigFile(path string) *Logger {
	config := loadConfigFile(path)
	return NewFileLoggerWithConfig(*config)
}

func NewFileLoggerWithConfig(cfg FileLoggerConfig) *Logger {
	logger := Logger{
		config: &cfg,
	}
	return &logger
}

func (l *Logger) GetConfig() FileLoggerConfig {
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

type levelLogger interface {
	Print(values map[string]interface{})
	Printf(format string, a ...interface{})
}

type emptyLevelLogger struct {
	levelLogger
}

func (*emptyLevelLogger) Print(values map[string]interface{})    {}
func (*emptyLevelLogger) Printf(format string, a ...interface{}) {}
