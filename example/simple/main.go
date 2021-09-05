package main

import (
	"github.com/snowlyg/iris-admin/server/web"
)

func main() {
	webServer := web.Init()
	webServer.Run()
}
