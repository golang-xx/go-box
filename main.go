package main

import (
	"flag"
	"fmt"
	"go-box/api"
	"go-box/common"
	"os"

	"github.com/labstack/echo"
)

const (
	version = "0.1"
)

var (
	h       bool
	v       bool
	cfgpath string
	e       *echo.Echo
)

func cmd() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&v, "v", false, "show version")
	flag.StringVar(&cfgpath, "c", "config.json", "set cfg and start")
	flag.Parse()

	if v {
		fmt.Println(version)
		os.Exit(0)
	}

	if h {
		flag.Usage()
		os.Exit(0)
	}
}

func main() {
	// cmd
	cmd()

	// config
	common.InitConfig(cfgpath)

	// api
	api.InitApi()

	// this will block forever
	api.StartAPP()
}