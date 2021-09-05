package main

import (
	"path/filepath"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/helper/dir"
	"github.com/snowlyg/iris-admin/server/web"
)

func main() {
	webServer := web.Init()
	webServer.AddStatic("/admin", iris.Dir(filepath.Join(dir.GetCurrentAbPath(), "admin")))
	webServer.Run()
}
