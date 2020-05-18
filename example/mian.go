package main

import (
	"fmt"

	"github.com/lostyear/glogger"
)

func main() {
	logger := glogger.NewLoggerWithConfigFile("config.toml")
	fmt.Println(logger.GetConfig())
}
