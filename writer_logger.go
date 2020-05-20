package glogger

import (
	"fmt"
	"io"
)

type writerLogger struct {
	*baseLogger
}

func NewWriterLoggerWithWriter(w io.Writer) ILogger {
	return NewWriterLoggerWithWriterAndConfig(w, LoggerConfig{})
}

func NewWriterLoggerWithWriterAndConfig(w io.Writer, cfg LoggerConfig) ILogger {
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
	// *log.Logger
	*baseLevelLogger
}

func newWriterLevelLogger(w io.Writer) *writerLevelLogger {
	return newWriterLevelLoggerWithConfig(w, LoggerConfig{})
}

func newWriterLevelLoggerWithConfig(w io.Writer, cfg LoggerConfig) *writerLevelLogger {
	cfg.validate()
	logger := newBaseLogger(w, fmt.Sprintf("[%s] ", cfg.Level), cfg.Flags)
	return &writerLevelLogger{
		baseLevelLogger: logger,
	}
}

// func (l *writerLevelLogger) Print(values map[string]interface{}) {
// 	sb := strings.Builder{}

// 	for k, v := range values {
// 		// TODO: field separator
// 		// TODO: file format
// 		if j, err := json.Marshal(v); err != nil {
// 			sb.WriteString(fmt.Sprintf("||%s=%v", k, v))
// 		} else {
// 			sb.WriteString(fmt.Sprintf("||%s=%s", k, string(j)))
// 		}
// 	}
// 	l.Logger.Println(sb.String())
// }

// func (l *writerLevelLogger) Printf(format string, a ...interface{}) {
// 	l.Logger.Printf(format, a...)
// }

type emptyWriter struct {
	io.Writer
}

func (ew *emptyWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}
