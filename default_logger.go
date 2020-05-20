package glogger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var defaultLogger ILogger

const (
	defaultLogFlags = log.Llongfile | log.Ldate | log.Ltime | log.Lmicroseconds
	defaultLogLevel = LLevelDebug
)

func init() {
	defaultLogger = NewWriterLoggerWithWriter(os.Stderr)
	// setDefaultLogger()
}

func setDefaultLogger() {
	// defaultLogger.debugLogger = newDefaultStdLogger(fmt.Sprintf("[%s] ", LLevelDebug))
}

func Debugf(format string, a ...interface{}) {
	defaultLogger.Debugf(format, a...)
	fmt.Println(defaultLogger)
}

type stdLogger struct {
	*log.Logger
}

func (l *stdLogger) Print(values map[string]interface{}) {
	sb := strings.Builder{}

	for k, v := range values {
		sb.WriteString(fmt.Sprintf("||%s=%v", k, v))
	}
	l.Logger.Println(sb.String())
}

func newDefaultStdLogger(prefix string) *stdLogger {
	l := log.New(log.Writer(), prefix, defaultLogFlags)
	return &stdLogger{
		Logger: l,
	}
}
