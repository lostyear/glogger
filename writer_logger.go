package glogger

import (
	"io"
)

type writerLogger struct {
	*baseLogger
}

var _ Logger = &writerLogger{}

func NewWriterLoggerWithWriter(w io.Writer) Logger {
	return NewWriterLoggerWithWriterAndConfig(w, LoggerConfig{})
}

func NewWriterLoggerWithWriterAndConfig(w io.Writer, cfg LoggerConfig) Logger {
	l := writerLogger{
		baseLogger: newBaseLoggerWithConfig(cfg),
	}

	for level, num := range lLevel {
		if num < lLevel[cfg.Level] {
			l.lLoggers[level] = defaultEmptyLogger
		} else {
			l.lLoggers[level] = newWriterLevelLoggerWithConfig(w, LoggerConfig{
				Flags: cfg.Flags,
				Level: level,
			})
		}
	}
	return &l
}

type writerLevelLogger struct {
	*baseLevelLogger
}

var _ levelLogger = &writerLevelLogger{}

func newWriterLevelLogger(w io.Writer) *writerLevelLogger {
	return newWriterLevelLoggerWithConfig(w, LoggerConfig{})
}

func newWriterLevelLoggerWithConfig(w io.Writer, cfg LoggerConfig) *writerLevelLogger {
	cfg.validate()
	logger := newBaseLogger(w, cfg.Level, cfg.Flags)
	return &writerLevelLogger{
		baseLevelLogger: logger,
	}
}

type emptyWriter struct{}

var _ io.Writer = &emptyWriter{}

func (ew *emptyWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}
