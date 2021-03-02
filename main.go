package main

import (
	"blumaton/bluock-core/config"
	"blumaton/bluock-core/scanner"
	"blumaton/bluock-core/utils"
	"flag"
	"runtime"
)

var (
	runtimeOS = runtime.GOOS
)

type flags struct {
	isConnectionTest bool
}

func parseArgs() flags {
	args := flags{}
	args.isConnectionTest = *flag.Bool("connection-test", false, "Start bluetooth connection test")

	return args
}

func main() {
	isSupportedOS()

	c, err := config.LoadConfig("config/example.yml")
	if err != nil {
		panic(err)
	}

	err = scanner.ScanDevices(&c)
	if err != nil {
		panic(err)
	}
}


func isSupportedOS() {
	switch runtimeOS {
	case "linux":
		return
	case "darwin":
		return
	default:
		utils.UnimplementedPanic()
	}
}
