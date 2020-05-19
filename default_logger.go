package glogger

import (
	"fmt"
	"log"
	"strings"
)

var defaultLogger Logger

const defaultLogFlags = log.Llongfile | log.Ldate | log.Ltime | log.Lmicroseconds

func init() {
	setDefaultLogger()
}

func setDefaultLogger() {
	defaultLogger.debugLogger = newDefaultStdLogger(fmt.Sprintf("[%s] ", LLevelDebug))
}

func Debugf(format string, a ...interface{}) {
	defaultLogger.Debugf(format, a...)
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
