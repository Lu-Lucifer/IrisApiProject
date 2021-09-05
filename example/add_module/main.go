package main

import (
	"fmt"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/iris-admin/g"
	"github.com/snowlyg/iris-admin/middleware"
	"github.com/snowlyg/iris-admin/server/database"
	"github.com/snowlyg/iris-admin/server/module"
	"github.com/snowlyg/iris-admin/server/validate"
	"github.com/snowlyg/iris-admin/server/web"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Party admin模块
func Party() module.WebModule {
	handler := func(admin iris.Party) {
		// 中间件
		admin.Use(middleware.InitCheck(), middleware.JwtHandler(), middleware.OperationRecord(), middleware.Casbin())
		admin.Get("/", GetAllAdmins).Name = "admin列表"
	}
	return module.NewModule("/admins", handler)
}

type ReqPaginate struct {
	g.Paginate
	Name string `json:"name"`
}

type Admin struct {
	gorm.Model
	Name string `gorm:"index;not null; type:varchar(60)" json:"name" `
}

func Paginate(db *gorm.DB, req ReqPaginate) (map[string]interface{}, error) {
	var count int64
	admins := []*Admin{}
	db = db.Model(&Admin{})
	if len(req.Name) > 0 {
		db = db.Where("name LIKE ?", fmt.Sprintf("%s%%", req.Name))
	}
	err := db.Count(&count).Error
	if err != nil {
		g.ZAPLOG.Error("获取用户总数错误", zap.String("错误:", err.Error()))
		return nil, err
	}

	err = db.Scopes(database.PaginateScope(req.Page, req.PageSize, req.Sort, req.OrderBy)).
		Find(&admins).Error
	if err != nil {
		g.ZAPLOG.Error("获取用户分页数据错误", zap.String("错误:", err.Error()))
		return nil, err
	}

	list := iris.Map{"items": admins, "total": count, "limit": req.PageSize}
	return list, nil
}

func GetAllAdmins(ctx iris.Context) {
	var req ReqPaginate
	if err := ctx.ReadJSON(&req); err != nil {
		errs := validate.ValidRequest(err)
		if len(errs) > 0 {
			g.ZAPLOG.Error("参数验证失败", zap.String("错误", strings.Join(errs, ";")))
			ctx.JSON(g.Response{Code: g.SystemErr.Code, Data: nil, Msg: strings.Join(errs, ";")})
			return
		}
	}

	list, err := Paginate(database.Instance(), req)
	if err != nil {
		ctx.JSON(g.Response{Code: g.SystemErr.Code, Data: nil, Msg: err.Error()})
		return
	}
	ctx.JSON(g.Response{Code: g.NoErr.Code, Data: list, Msg: g.NoErr.Msg})
}

func main() {
	webServer := web.Init()
	webServer.AddModule(Party())
	webServer.Run()
}
