package main

import (
	v1 "github.com/snowlyg/iris-admin/modules/v1"
	"github.com/snowlyg/iris-admin/server/web"
)

func main() {
	webServer := web.Init()
	// 添加模块
	webServer.AddModule(v1.Party())
	// 启动
	webServer.Run()
}
