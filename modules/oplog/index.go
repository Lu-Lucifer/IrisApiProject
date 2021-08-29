package oplog

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/iris-admin/middleware"
	"github.com/snowlyg/iris-admin/server/module"
)

// Party 调试模块
func Party() module.WebModule {
	handler := func(index iris.Party) {
		index.Use(middleware.InitCheck())
	}
	return module.NewModule("/oplog", handler)
}
