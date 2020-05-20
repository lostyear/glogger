package glogger

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type WriterLogger struct {
	ILogger

	config   LoggerConfig
	lLoggers map[string]levelLogger
}

func NewWriterLoggerWithWriter(w io.Writer) ILogger {
	return NewWriterLoggerWithWriterAndConfig(w, LoggerConfig{})
}

func NewWriterLoggerWithWriterAndConfig(w io.Writer, cfg LoggerConfig) ILogger {
	if cfg.Flags == 0 {
		cfg.Flags = defaultLogFlags
	}
	ulevel := strings.ToUpper(cfg.Level)
	if _, ok := lLevel[ulevel]; !ok {
		ulevel = defaultLogLevel
	}
	cfg.Level = ulevel

	l := WriterLogger{
		config: cfg,
	}

	l.lLoggers = make(map[string]levelLogger)
	for level, num := range lLevel {
		if num < lLevel[cfg.Level] {

			l.lLoggers[level] = &emptyLevelLogger{}
		} else {
			l.lLoggers[level] = newWriterLevelLoggerWithConfig(w, LoggerConfig{
				Flags: cfg.Flags,
				Level: cfg.Level,
			})
		}
	}
	return &l
}

func (l *WriterLogger) Debug(values map[string]interface{}) {
	l.lLoggers[LLevelDebug].Print(values)
}
func (l *WriterLogger) Debugf(format string, a ...interface{}) {
	l.lLoggers[LLevelDebug].Printf(format, a...)
}
func (l *WriterLogger) Info(values map[string]interface{}) {
	l.lLoggers[LLevelInfo].Print(values)
}
func (l *WriterLogger) Infof(format string, a ...interface{}) {
	l.lLoggers[LLevelInfo].Printf(format, a...)
}
func (l *WriterLogger) Warn(values map[string]interface{}) {
	l.lLoggers[LLevelWarn].Print(values)
}
func (l *WriterLogger) Warnf(format string, a ...interface{}) {
	l.lLoggers[LLevelWarn].Printf(format, a...)
}
func (l *WriterLogger) Error(values map[string]interface{}) {
	l.lLoggers[LLevelError].Print(values)
}
func (l *WriterLogger) Errorf(format string, a ...interface{}) {
	l.lLoggers[LLevelError].Printf(format, a...)
}
func (l *WriterLogger) Fatal(values map[string]interface{}) {
	l.lLoggers[LLevelFatal].Print(values)
	os.Exit(1)
}
func (l *WriterLogger) Fatalf(format string, a ...interface{}) {
	l.lLoggers[LLevelFatal].Printf(format, a...)
	os.Exit(1)
}

type writerLevelLogger struct {
	levelLogger

	*log.Logger
}

func newWriterLevelLogger(w io.Writer) *writerLevelLogger {
	return newWriterLevelLoggerWithConfig(w, LoggerConfig{})
}

func newWriterLevelLoggerWithConfig(w io.Writer, cfg LoggerConfig) *writerLevelLogger {
	var flags int
	var level string

	if cfg.Flags == 0 {
		flags = defaultLogFlags
	} else {
		flags = cfg.Flags
	}

	level = strings.ToUpper(cfg.Level)
	if _, ok := lLevel[level]; !ok {
		level = defaultLogLevel
	}

	logger := log.New(w, fmt.Sprintf("[%s] ", level), flags)
	return &writerLevelLogger{
		Logger: logger,
	}
}

func (l *writerLevelLogger) Print(values map[string]interface{}) {
	sb := strings.Builder{}

	for k, v := range values {
		if j, err := json.Marshal(v); err != nil {
			sb.WriteString(fmt.Sprintf("||%s=%v", k, v))
		} else {
			sb.WriteString(fmt.Sprintf("||%s=%s", k, string(j)))
		}
	}
	l.Logger.Println(sb.String())
}

func (l *writerLevelLogger) Printf(format string, a ...interface{}) {
	l.Logger.Printf(format, a...)
}

type emptyWriter struct {
	io.Writer
}

func (ew *emptyWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}
