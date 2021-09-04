package main

import (
	"github.com/snowlyg/iris-admin/server/web"
)

var Version = "2.0"

func main() {
	webServer := web.Init()
	webServer.Run()
}
