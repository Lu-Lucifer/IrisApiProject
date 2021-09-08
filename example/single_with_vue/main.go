package main

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/iris-admin/server/web"
)

func main() {
	webServer := web.Init()
	webServer.AddStatic("/", iris.Dir("./dist"), iris.DirOptions{
		IndexName: "index.html",
		SPA:       true,
	})
	webServer.Run()
}
