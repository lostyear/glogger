package glogger

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type ILogger interface {
	IConfig

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

func newBaseLoggerWithConfig(cfg LoggerConfig) *baseLogger {
	cfg.validate()
	return &baseLogger{
		config:   cfg,
		lLoggers: make(mLogger),
	}
}

func (l *baseLogger) validate() {
	l.config.validate()
}

func (l *baseLogger) GetConfig() IConfig {
	return &l.config
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

type baseLevelLogger struct {
	*log.Logger
}

func newBaseLogger(w io.Writer, level string, flag int) *baseLevelLogger {
	return &baseLevelLogger{
		Logger: log.New(w, fmt.Sprintf("[%s] ", level), flag),
	}
}

func (l *baseLevelLogger) Print(values map[string]interface{}) {
	sb := strings.Builder{}

	for k, v := range values {
		// TODO: field separator
		// TODO: file format
		if j, err := json.Marshal(v); err != nil {
			sb.WriteString(fmt.Sprintf("||%s=%v", k, v))
		} else {
			sb.WriteString(fmt.Sprintf("||%s=%s", k, string(j)))
		}
	}
	l.Logger.Println(sb.String())
}

func (l *baseLevelLogger) Printf(format string, a ...interface{}) {
	l.Logger.Printf(format, a...)
}

type emptyLevelLogger struct{}

var defaultEmptyLogger levelLogger = &emptyLevelLogger{}

func (*emptyLevelLogger) Print(values map[string]interface{})    {}
func (*emptyLevelLogger) Printf(format string, a ...interface{}) {}
