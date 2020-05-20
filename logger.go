package glogger

import "os"

type ILogger interface {
	GetConfig() LoggerConfig
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

type baseLogger struct {
	config   LoggerConfig
	lLoggers mLogger
}

var _ ILogger = &baseLogger{}

func (l *baseLogger) GetConfig() LoggerConfig {
	return l.config
}

func (l *baseLogger) Debug(values map[string]interface{}) {
	l.lLoggers[LLevelDebug].Print(values)
}
func (l *baseLogger) Debugf(format string, a ...interface{}) {
	l.lLoggers[LLevelDebug].Printf(format, a...)
}
func (l *baseLogger) Info(values map[string]interface{}) {
	l.lLoggers[LLevelInfo].Print(values)
}
func (l *baseLogger) Infof(format string, a ...interface{}) {
	l.lLoggers[LLevelInfo].Printf(format, a...)
}
func (l *baseLogger) Warn(values map[string]interface{}) {
	l.lLoggers[LLevelWarn].Print(values)
}
func (l *baseLogger) Warnf(format string, a ...interface{}) {
	l.lLoggers[LLevelWarn].Printf(format, a...)
}
func (l *baseLogger) Error(values map[string]interface{}) {
	l.lLoggers[LLevelError].Print(values)
}
func (l *baseLogger) Errorf(format string, a ...interface{}) {
	l.lLoggers[LLevelError].Printf(format, a...)
}
func (l *baseLogger) Fatal(values map[string]interface{}) {
	l.lLoggers[LLevelFatal].Print(values)
	os.Exit(1)
}
func (l *baseLogger) Fatalf(format string, a ...interface{}) {
	l.lLoggers[LLevelFatal].Printf(format, a...)
	os.Exit(1)
}

type mLogger map[string]levelLogger

type levelLogger interface {
	Print(values map[string]interface{})
	Printf(format string, a ...interface{})
}

type emptyLevelLogger struct{}

var defaultEmptyLogger levelLogger = &emptyLevelLogger{}

func (*emptyLevelLogger) Print(values map[string]interface{})    {}
func (*emptyLevelLogger) Printf(format string, a ...interface{}) {}
