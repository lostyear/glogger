package glogger

const (
	LLevelDebug = "DEBUG"
	LLevelInfo  = "INFO"
	LLevelWarn  = "WARN"
	LLevelError = "ERROR"
	LLevelFatal = "FATAL"
	LLevelOff   = "OFF"
)

const (
	lLevelDebugNum = iota
	lLevelInfoNum
	lLevelWarnNum
	lLevelErrorNum
	lLevelFatalNum
	lLevelOffNum
)

var lLevel = map[string]int{
	LLevelDebug: lLevelDebugNum,
	LLevelInfo:  lLevelInfoNum,
	LLevelWarn:  lLevelWarnNum,
	LLevelError: lLevelErrorNum,
	LLevelFatal: lLevelFatalNum,
	LLevelOff:   lLevelOffNum,
}
