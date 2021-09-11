package main

import (
	"path/filepath"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/helper/dir"
	v1 "github.com/snowlyg/iris-admin/modules/v1"
	"github.com/snowlyg/iris-admin/server/web"
)

func main() {
	webServer := web.Init()
	// 添加前端页面
	webServer.AddStatic("/", iris.Dir("./dist"), iris.DirOptions{
		IndexName: "index.html",
		SPA:       true,
	})
	// 添加上传文件路径
	webServer.AddStatic("/upload", iris.Dir(filepath.Join(dir.GetCurrentAbPath(), "/static/upload")))
	// 添加模块
	webServer.AddModule(v1.Party())
	// 启动
	webServer.Run()
}
