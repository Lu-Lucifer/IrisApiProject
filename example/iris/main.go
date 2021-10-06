package main

import (
	"github.com/snowlyg/iris-admin/server/web"
	"github.com/snowlyg/iris-admin/server/web/web_iris"
)

func main() {
	web.Start(web_iris.Init())
}
