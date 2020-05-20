package glogger

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

type writerLogger struct {
	*baseLogger
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

	l := writerLogger{
		baseLogger: &baseLogger{
			config: cfg,
		},
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

type writerLevelLogger struct {
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
