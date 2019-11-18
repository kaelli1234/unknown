package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"

	"ccgwf/engine"

	"unknown/app"
	_ "unknown/routers"
)

var (
	BuildUser    = ""
	BuildTime    = ""
	BuildGcc     = ""
	BuildVersion = ""
	BuildMachine = ""

	configPath = flag.String("f", "", "config file path")
	version    = flag.Bool("v", false, "version")
)

func init() {
	flag.Parse() // init args
	if *version {
		fmt.Printf("VERSION    : %s\n", strings.Replace(BuildVersion, "_", " ", -1))
		fmt.Printf("BUILD BY   : %s\n", strings.Replace(BuildUser, "_", " ", -1))
		fmt.Printf("BUILD TIME : %s\n", strings.Replace(BuildTime, "_", " ", -1))
		fmt.Printf("ON MACHINE : %s\n", strings.Replace(BuildMachine, "_", " ", -1))
		fmt.Printf("USE GCC    : %s\n", strings.Replace(BuildGcc, "_", " ", -1))
		os.Exit(0)
	}
	if len(*configPath) == 0 {
		fmt.Printf("-f miss")
		os.Exit(0)
	}
}

func main() {
	app.Init(configPath)
	runtime.GOMAXPROCS(runtime.NumCPU())
	engine.Run()
}
