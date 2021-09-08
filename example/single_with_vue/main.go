package main

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/iris-admin/server/web"
)

func main() {
	webServer := web.Init()
	templatesFS := iris.PrefixDir("./dist", AssetFile())
	publicFS := iris.PrefixDir("./dist", AssetFile())
	webServer.AddStatic(templatesFS, publicFS)
	webServer.Run()
}
