package main

import (
	"fmt"

	"github.com/lostyear/glogger"
)

func main() {
	logger := glogger.NewFileLoggerWithConfigFile("config.toml")
	fmt.Println(logger.GetConfig())
	glogger.Debugf("This is a debug log")
}
