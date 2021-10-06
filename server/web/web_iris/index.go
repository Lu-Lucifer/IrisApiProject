package web_iris

import (
	stdContext "context"
	"errors"
	"fmt"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/helper/dir"
	"github.com/snowlyg/helper/str"
	"github.com/snowlyg/helper/tests"
	"github.com/snowlyg/iris-admin/server/cache"
	"github.com/snowlyg/iris-admin/server/viper_server"
	"github.com/snowlyg/multi"
)

var ErrAuthDriverEmpty = errors.New("认证驱动初始化失败")

// WebServer web服务
// - app iris application
// - idleConnsClosed
// - addr  服务访问地址
// - timeFormat  时间格式
// - staticPrefix  静态文件访问地址前缀
// - staticPath  静态文件地址
// - webPath  前端文件地址
type WebServer struct {
	app             *iris.Application
	idleConnsClosed chan struct{}
	addr            string
	timeFormat      string
	staticPrefix    string
	staticPath      string
	webPrefix       string
	webPath         string
}

// Init 初始化web服务
// 先初始化基础服务 config , zap , database , casbin  e.g.
func Init() *WebServer {
	viper_server.Init(getViperConfig())
	app := iris.New()
	app.Validator = validator.New() //参数验证
	app.Logger().SetLevel(CONFIG.System.Level)
	idleConnsClosed := make(chan struct{})
	iris.RegisterOnInterrupt(func() { //优雅退出
		timeout := 10 * time.Second
		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
		defer cancel()
		app.Shutdown(ctx) // close all hosts
		close(idleConnsClosed)
	})

	if CONFIG.System.Addr == "" { // 默认 8085
		CONFIG.System.Addr = "127.0.0.1:8085"
	}

	if CONFIG.System.StaticPath == "" { // 默认 /static/upload
		CONFIG.System.StaticPath = "/static/upload"
	}

	if CONFIG.System.StaticPrefix == "" { // 默认 /upload
		CONFIG.System.StaticPrefix = "/upload"
	}

	if CONFIG.System.WebPath == "" { // 默认 ./dist
		CONFIG.System.WebPath = "./dist"
	}

	if CONFIG.System.WebPrefix == "" { // 默认 /
		CONFIG.System.WebPrefix = "/"
	}

	if CONFIG.System.TimeFormat == "" { // 默认 80
		CONFIG.System.TimeFormat = "2006-01-02 15:04:05"
	}

	return &WebServer{
		app:             app,
		addr:            CONFIG.System.Addr,
		timeFormat:      CONFIG.System.TimeFormat,
		staticPrefix:    CONFIG.System.StaticPrefix,
		staticPath:      CONFIG.System.StaticPath,
		webPrefix:       CONFIG.System.WebPrefix,
		webPath:         CONFIG.System.WebPath,
		idleConnsClosed: idleConnsClosed,
	}
}

// AddStatic 添加静态文件
func (ws *WebServer) AddStatic(requestPath string, fsOrDir interface{}, opts ...iris.DirOptions) {
	ws.app.HandleDir(requestPath, fsOrDir, opts...)
}

// InitDriver 初始化认证
func (ws *WebServer) InitDriver() error {
	err := multi.InitDriver(
		&multi.Config{
			DriverType:      CONFIG.System.CacheType,
			UniversalClient: cache.Instance()},
	)
	if err != nil {
		return fmt.Errorf("初始化认证驱动错误 %w", err)
	}
	if multi.AuthDriver == nil {
		return ErrAuthDriverEmpty
	}
	return nil
}

// AddWebStatic 添加前端访问地址
func (ws *WebServer) AddWebStatic(perfix string) {
	fsOrDir := iris.Dir(filepath.Join(dir.GetCurrentAbPath(), ws.webPath))
	opt := iris.DirOptions{
		IndexName: "index.html",
		SPA:       true,
	}
	ws.AddStatic(perfix, fsOrDir, opt)
}

// AddUploadStatic 添加上传文件访问地址
func (ws *WebServer) AddUploadStatic() {
	fsOrDir := iris.Dir(filepath.Join(dir.GetCurrentAbPath(), ws.staticPath))
	ws.AddStatic(ws.staticPrefix, fsOrDir)
}

// GetTestClient 获取测试验证客户端
func (ws *WebServer) GetTestClient(t *testing.T) *tests.Client {
	var once sync.Once
	var client *tests.Client
	once.Do(
		func() {
			client = tests.New(str.Join("http://", ws.addr), t, ws.app)
			if client == nil {
				t.Errorf("test client is nil")
			}
		},
	)

	return client
}

// GetTestLogin 测试登录web服务
func (ws *WebServer) GetTestLogin(t *testing.T, url string, res tests.Responses, datas ...map[string]interface{}) *tests.Client {
	client := ws.GetTestClient(t)
	if client == nil {
		return nil
	}
	err := client.Login(url, res, datas...)
	if err != nil {
		t.Error(err)
		return nil
	}
	return client
}

// Run 启动web服务
func (ws *WebServer) Run() {
	err := ws.InitRouter()
	if err != nil {
		fmt.Printf("初始化路由错误： %v\n", err)
		panic(err)
	}
	ws.app.Listen(
		ws.addr,
		iris.WithoutInterruptHandler,
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		iris.WithTimeFormat(ws.timeFormat),
	)
	<-ws.idleConnsClosed
}
