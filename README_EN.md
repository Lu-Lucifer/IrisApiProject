<h1 align="center">IrisAdmin</h1>

[![Build Status](https://app.travis-ci.com/snowlyg/iris-admin.svg?branch=master)](https://app.travis-ci.com/snowlyg/iris-admin)
[![LICENSE](https://img.shields.io/github/license/snowlyg/iris-admin)](https://github.com/snowlyg/iris-admin/blob/master/LICENSE)
[![go doc](https://godoc.org/github.com/snowlyg/iris-admin?status.svg)](https://godoc.org/github.com/snowlyg/iris-admin)
[![go report](https://goreportcard.com/badge/github.com/snowlyg/iris-admin)](https://goreportcard.com/badge/github.com/snowlyg/iris-admin)
[![Build Status](https://codecov.io/gh/snowlyg/iris-admin/branch/master/graph/badge.svg)](https://codecov.io/gh/snowlyg/iris-admin)

[简体中文](./README.md)  | English

#### Project url

[GITHUB](https://github.com/snowlyg/iris-admin) | [GITEE](https://gitee.com/snowlyg/iris-admin)
****
> This project just for learning golang, welcome to give your suggestions!

#### Documentation

- [IRIS-ADMIN-DOC](https://doc.snowlyg.com)
- [IRIS V12 document for chinese](https://github.com/snowlyg/iris/wiki)
- [godoc](https://pkg.go.dev/github.com/snowlyg/iris-admin?utm_source=godoc)

[![Gitter](https://badges.gitter.im/iris-go-tenancy/community.svg)](https://gitter.im/iris-go-tenancy/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge) [![Join the chat at https://gitter.im/iris-go-tenancy/iris-admin](https://badges.gitter.im/iris-go-tenancy/iris-admin.svg)](https://gitter.im/iris-go-tenancy/iris-admin?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
#### BLOG

- [REST API with iris-go web framework](https://blog.snowlyg.com/iris-go-api-1/)

- [How to user iris-go with casbin](https://blog.snowlyg.com/iris-go-api-2/)

---

#### Getting started

- Get master package , Notice must use `master` version.

```sh
 go get github.com/snowlyg/iris-admin@master
```

#### Program introduction

##### The project consists of multiple plugins, each with different functions

- [viper_server]
  - The plugin configuration is initialized and generate a local configuration file.
  - Use [github.com/spf13/viper](https://github.com/spf13/viper) third party package.
  - Need implement  `func getViperConfig() viper_server.ViperConfig` function.

```go
package cache

import (
  "fmt"

  "github.com/fsnotify/fsnotify"
  "github.com/snowlyg/iris-admin/g"
  "github.com/snowlyg/iris-admin/server/viper_server"
  "github.com/spf13/viper"
)

var CONFIG Redis

type Redis struct {
  DB       int    `mapstructure:"db" json:"db" yaml:"db"`
  Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
  Password string `mapstructure:"password" json:"password" yaml:"password"`
  PoolSize int    `mapstructure:"pool-size" json:"poolSize" yaml:"pool-size"`
}

// getViperConfig get initialize config
func getViperConfig() viper_server.ViperConfig {
  configName := "redis"
  db := fmt.Sprintf("%d", CONFIG.DB)
  poolSize := fmt.Sprintf("%d", CONFIG.PoolSize)
  return viper_server.ViperConfig{
    Directory: g.ConfigDir,
    Name:      configName,
    Type:      g.ConfigType,
    Watch: func(vi *viper.Viper) error {
      if err := vi.Unmarshal(&CONFIG); err != nil {
        return fmt.Errorf("deserialization data error: %v", err)
      }
      // config file change
      vi.SetConfigName(configName)
      return nil
    },
    // Note: When setting the default configuration value, there can be no other symbols such as spaces in front. It must be close to the left
    Default: []byte(`
db: ` + db + `
addr: "` + CONFIG.Addr + `"
password: "` + CONFIG.Password + `"
pool-size: ` + poolSize),
  }
}
```

- [zap_server]
  - Plugin logging.
  - Use [go.uber.org/zap](https://pkg.go.dev/go.uber.org/zap) third party package.
  - Through global variables `zap_server.ZAPLOG` record the log of the corresponding level.
  
```go
  zap_server.ZAPLOG.Info("Registration data table error", zap.Any("err", err))
  zap_server.ZAPLOG.Debug("Registration data table error", zap.Any("err", err))
  zap_server.ZAPLOG.Error("Registration data table error", zap.Any("err", err))
  ...
```

- [database]
  - database plugin [only support mysql now].
  - Use [gorm.io/gorm](https://github.com/go-gorm/gorm) third party package.
  - Through single instance `database.Instance()` operating data.
  
```go
  database.Instance().Model(&User{}).Where("name = ?","name").Find(&user)
  ...
```

- [casbin]
  - Access control management plugin.
  - Use [casbin](github.com/casbin/casbin/v2 ) third party package.
  - Through use `casbin.Instance()` middleware on route,implement interface authority authentication

```go
	_, err := casbin.Instance().AddRoleForUser("1", "999") 
	uids, err := casbin.Instance().GetRolesForUser("1") 
	_, err := casbin.Instance().RemoveFilteredPolicy(v, p...) 
  ...
```

- [cache]
  - Cache-driven plugin
  - Use [github.com/go-redis/redis](https://github.com/go-redis/redis) third party package.
  - Through single instance `cache.Instance()` operating data.

```go
  	err := cache.Instance().Set(context.Background(), "key", "value", expiration).Err()
    cache.Instance().Del(context.Background(), "key").Result()
    cache.Instance().Get(context.Background(), "key")
  ...
```

- [operation]
  - System operation log plugin.
  - Through use `index.Use(operation.OperationRecord())` middleware on route , realize the interface to automatically generate operation logs.

- [cron_server]
  - Job server
  - Use [robfig/cron](https://github.com/robfig/cron) third party package.
  - Through single instance `cron_server.Instance()` to add job or func.
  
```go
  cron_server.CronInstance().AddJob("@every 1m",YourJob)
  // or 
  cron_server.CronInstance().AddFunc("@every 1m",YourFunc)
  ...
```

- [web]
  - web_iris [Go-Iris](https://github.com/kataras/iris) web framework plugin.
  - web_gin [Go-gin web](https://github.com/gin-gonic/gin) web framework plugin.
  - web framework plugin need implement `type WebFunc interface {}`  interface.

-

```go
type WebBaseFunc interface {
  AddWebStatic(staticAbsPath, webPrefix string, paths ...string)
  AddUploadStatic(staticAbsPath, webPrefix string)
  InitRouter() error
  Run()
}

// WebFunc web framework
// - GetTestClient test client
// - GetTestLogin test for login
// - AddWebStatic add web static path
// - AddUploadStatic add upload static path 
// - Run start
type WebFunc interface {
  WebBaseFunc
}
```

- [mongodb]
  - mongodb
  - Use [mongodb](https://www.mongodb.com/) third party package.

#### Initialize database

##### Simple

- Use gorm's `AutoMigrate()` function to auto migrate database.

```go
package main

import (
  "github.com/snowlyg/iris-admin/server/web"
  "github.com/snowlyg/iris-admin/server/web/web_iris"
  "github.com/snowlyg/iris-admin-rbac/iris/perm"
  "github.com/snowlyg/iris-admin-rbac/iris/role"
  "github.com/snowlyg/iris-admin/server/database"
  "github.com/snowlyg/iris-admin/server/operation"
)

func main() {
    database.Instance().AutoMigrate(&perm.Permission{},&role.Role{},&user.User{},&operation.Oplog{})
}
```

##### Custom migrate tools

- Use `gormigrate` third party package. Tt's helpful for database migrate and program development.
- Detail is see  [iris-admin-cmd](https://github.com/snowlyg/iris-admin-example/blob/main/iris/cmd/main.go).
  
---

- Add main.go file.

```go
package main

import (
  "github.com/snowlyg/iris-admin/server/web"
  "github.com/snowlyg/iris-admin/server/web/web_iris"
)

func main() {
  wi := web_iris.Init()
  web.Start(wi)
}
```

#### Run project

- When you first run this cmd `go run main.go` , you can see some config files in  the `config` directory,
- and `rbac_model.conf` will be created in your project root directory.

```sh
go run main.go
```

#### Module

- You can use [iris-admin-rbac](https://github.com/snowlyg/iris-admin-rbac) package to add rbac function for your project quickly.
- Your can use AddModule() to add other modules .

```go
package main

import (
  rbac "github.com/snowlyg/iris-admin-rbac/iris"
  "github.com/snowlyg/iris-admin/server/web"
  "github.com/snowlyg/iris-admin/server/web/web_iris"
)

func main() {
  wi := web_iris.Init()
  rbacParty := web_iris.Party{
    Perfix:    "/api/v1",
    PartyFunc: rbac.Party(),
  }
  wi.AddModule(rbacParty)
  web.Start(web_iris.Init())
}
```

#### Default static file path

- A static file access path has been built in by default
- Static files will upload to `/static/upload` directory.
- You can set this config key `static-path` to change the default directory.

```yaml
system:
  addr: "127.0.0.1:8085"
  db-type: ""
  level: debug
  static-prefix: /upload
  time-format: "2006-01-02 15:04:05"
  web-path: ./dist
```

#### Use with front-end framework , e.g. vue

- Default,you must build vue to the `dist` directory.
- Naturally you can set this config key `web-path` to change the default directory.
  
```go
package main

import (
  "github.com/kataras/iris/v12"
  "github.com/snowlyg/iris-admin/server/web"
)

func main() {
  webServer := web_iris.Init()
  wi.AddUploadStatic("/upload", "/var/static")
  wi.AddWebStatic("/", "/var/static")
  webServer.Run()
}
```

- Front-end page reference/borrowing:
 *notice: The front-end only realizes preview effect simply*
- [gin-vue-admin](https://github.com/flipped-aurora/gin-vue-admin/tree/master/web)
- [vue-element-admin](https://github.com/PanJiaChen/vue-element-admin)

#### Example

- [iris](https://github.com/snowlyg/iris-admin-example/tree/main/iris)
- [gin](https://github.com/snowlyg/iris-admin-example/tree/main/gin)

#### RBAC

- [iris-admin-rbac](https://github.com/snowlyg/iris-admin-rbac)

#### Unit test and documentation

- Before start unit tests, you need to set two system environment variables `mysqlPwd` and `mysqlAddr`,that will be used when running the test instance。
- [helper/tests](https://github.com/snowlyg/helper/tree/main/tests) package the unit test used, it's  simple package base on [httpexpect/v2](https://github.com/gavv/httpexpect).
- [example for unit test](https://github.com/snowlyg/iris-admin-rbac/tree/main/iris/perm/tests)
- [example for unit test](https://github.com/snowlyg/iris-admin-rbac/tree/main/gin/authority/test)

Before create a http api unit test , you need create a base test file named `main_test.go` , this file have some unit test step ：
***Suggest use docker mysql, otherwise if the test fails, there will be a lot of test data left behind***
- 1.create database before test start and delete database when test finish.
- 2.create tables and seed test data at once time.
- 3.`PartyFunc` and `SeedFunc` use to custom someting for your test model.
内容如下所示:
***main_test.go***

```go
package test

import (
  "os"
  "testing"

  "github.com/snowlyg/httptest"
  rbac "github.com/snowlyg/iris-admin-rbac/gin"
  "github.com/snowlyg/iris-admin/server/web/common"
  "github.com/snowlyg/iris-admin/server/web/web_gin"
)

var TestServer *web_gin.WebServer
var TestClient *httptest.Client

func TestMain(m *testing.M) {
  var uuid string
  uuid, TestServer = common.BeforeTestMainGin(rbac.PartyFunc, rbac.SeedFunc)
  code := m.Run()
  common.AfterTestMain(uuid, true)

  os.Exit(code)
}

```

***index_test.go***

```go
package test

import (
  "fmt"
  "net/http"
  "path/filepath"
  "testing"

  "github.com/snowlyg/helper/str"
  "github.com/snowlyg/httptest"
  rbac "github.com/snowlyg/iris-admin-rbac/gin"
  "github.com/snowlyg/iris-admin/g"
  "github.com/snowlyg/iris-admin/server/web"
  "github.com/snowlyg/iris-admin/server/web/web_gin/response"
)

var (
  url = "/api/v1/admin"
)

func TestList(t *testing.T) {
  TestClient = httptest.Instance(t, str.Join("http://", web.CONFIG.System.Addr), TestServer.GetEngine())
  TestClient.Login(rbac.LoginUrl, nil)
  if TestClient == nil {
    return
  }
  pageKeys := httptest.Responses{
    {Key: "status", Value: http.StatusOK},
    {Key: "message", Value: response.ResponseOkMessage},
    {Key: "data", Value: httptest.Responses{
      {Key: "pageSize", Value: 10},
      {Key: "page", Value: 1},
      {Key: "list", Value: []httptest.Responses{
        {
          {Key: "id", Value: 1, Type: "ge"},
          {Key: "nickName", Value: "superadmin"},
          {Key: "username", Value: "admin"},
          {Key: "headerImg", Value: "http://xxxx/head.png"},
          {Key: "status", Value: g.StatusTrue},
          {Key: "isShow", Value: g.StatusFalse},
          {Key: "phone", Value: "13800138000"},
          {Key: "email", Value: "admin@admin.com"},
          {Key: "authorities", Value: []string{"superadmin"}},
          {Key: "updatedAt", Value: "", Type: "notempty"},
          {Key: "createdAt", Value: "", Type: "notempty"},
        },
      }},
      {Key: "total", Value: 0, Type: "ge"},
    }},
  }
  TestClient.GET(fmt.Sprintf("%s/getAll", url), pageKeys, httptest.RequestParams)
}

func TestCreate(t *testing.T) {
  TestClient = httptest.Instance(t, str.Join("http://", web.CONFIG.System.Addr), TestServer.GetEngine())
  TestClient.Login(rbac.LoginUrl, nil)
  if TestClient == nil {
    return
  }

  data := map[string]interface{}{
    "nickName":     "test name",
    "username":     "create_test_username",
    "authorityIds": []uint{web.AdminAuthorityId},
    "email":        "get@admin.com",
    "phone":        "13800138001",
    "password":     "123456",
  }
  id := Create(TestClient, data)
  if id == 0 {
    t.Fatalf("add user failed by id=%d", id)
  }
  defer Delete(TestClient, id)
}

func TestUpdate(t *testing.T) {
  TestClient = httptest.Instance(t, str.Join("http://", web.CONFIG.System.Addr), TestServer.GetEngine())
  TestClient.Login(rbac.LoginUrl, nil)
  if TestClient == nil {
    return
  }
  data := map[string]interface{}{
    "nickName":     "test name",
    "username":     "create_test_username_for_update",
    "authorityIds": []uint{web.AdminAuthorityId},
    "email":        "get@admin.com",
    "phone":        "13800138001",
    "password":     "123456",
  }
  id := Create(TestClient, data)
  if id == 0 {
    t.Fatalf("add user failed by id=%d", id)
  }
  defer Delete(TestClient, id)

  update := map[string]interface{}{
    "nickName": "test name",
    "email":    "get@admin.com",
    "phone":    "13800138003",
    "password": "123456",
  }

  pageKeys := httptest.Responses{
    {Key: "status", Value: http.StatusOK},
    {Key: "message", Value: response.ResponseOkMessage},
  }
  TestClient.PUT(fmt.Sprintf("%s/updateAdmin/%d", url, id), pageKeys, update)
}

func TestGetById(t *testing.T) {
  TestClient = httptest.Instance(t, str.Join("http://", web.CONFIG.System.Addr), TestServer.GetEngine())
  TestClient.Login(rbac.LoginUrl, nil)
  if TestClient == nil {
    return
  }
  data := map[string]interface{}{
    "nickName":     "test name",
    "username":     "create_test_username_for_get",
    "email":        "get@admin.com",
    "phone":        "13800138001",
    "authorityIds": []uint{web.AdminAuthorityId},
    "password":     "123456",
  }
  id := Create(TestClient, data)
  if id == 0 {
    t.Fatalf("add user failed by id=%d", id)
  }
  defer Delete(TestClient, id)
  pageKeys := httptest.Responses{
    {Key: "status", Value: http.StatusOK},
    {Key: "message", Value: response.ResponseOkMessage},
    {Key: "data", Value: httptest.Responses{
      {Key: "id", Value: 1, Type: "ge"},
      {Key: "nickName", Value: data["nickName"].(string)},
      {Key: "username", Value: data["username"].(string)},
      {Key: "status", Value: g.StatusTrue},
      {Key: "email", Value: data["email"].(string)},
      {Key: "phone", Value: data["phone"].(string)},
      {Key: "isShow", Value: g.StatusTrue},
      {Key: "headerImg", Value: "http://xxxx/head.png"},
      {Key: "updatedAt", Value: "", Type: "notempty"},
      {Key: "createdAt", Value: "", Type: "notempty"},
      {Key: "createdAt", Value: "", Type: "notempty"},
      {Key: "authorities", Value: []string{"超级管理员"}},
    },
    },
  }
  TestClient.GET(fmt.Sprintf("%s/getAdmin/%d", url, id), pageKeys)
}

func TestChangeAvatar(t *testing.T) {
  TestClient = httptest.Instance(t, str.Join("http://", web.CONFIG.System.Addr), TestServer.GetEngine())
  TestClient.Login(rbac.LoginUrl, nil)
  if TestClient == nil {
    return
  }
  data := map[string]interface{}{
    "headerImg": "/avatar.png",
  }
  pageKeys := httptest.Responses{
    {Key: "status", Value: http.StatusOK},
    {Key: "message", Value: response.ResponseOkMessage},
  }
  TestClient.POST(fmt.Sprintf("%s/changeAvatar", url), pageKeys, data)

  profile := httptest.Responses{
    {Key: "status", Value: http.StatusOK},
    {Key: "message", Value: response.ResponseOkMessage},
    {Key: "data", Value: httptest.Responses{
      {Key: "id", Value: 1, Type: "ge"},
      {Key: "nickName", Value: "superadmin"},
      {Key: "username", Value: "admin"},
      {Key: "headerImg", Value: filepath.ToSlash(web.ToStaticUrl("/avatar.png"))},
      {Key: "status", Value: g.StatusTrue},
      {Key: "isShow", Value: g.StatusFalse},
      {Key: "phone", Value: "13800138000"},
      {Key: "email", Value: "admin@admin.com"},
      {Key: "authorities", Value: []string{"superadmin"}},
      {Key: "updatedAt", Value: "", Type: "notempty"},
      {Key: "createdAt", Value: "", Type: "notempty"},
    },
    },
  }
  TestClient.GET(fmt.Sprintf("%s/profile", url), profile)
}

func Create(TestClient *httptest.Client, data map[string]interface{}) uint {
  pageKeys := httptest.Responses{
    {Key: "status", Value: http.StatusOK},
    {Key: "message", Value: response.ResponseOkMessage},
    {Key: "data", Value: httptest.Responses{
      {Key: "id", Value: 1, Type: "ge"},
    },
    },
  }
  return TestClient.POST(fmt.Sprintf("%s/createAdmin", url), pageKeys, data).GetId()
}

func Delete(TestClient *httptest.Client, id uint) {
  pageKeys := httptest.Responses{
    {Key: "status", Value: http.StatusOK},
    {Key: "message", Value: response.ResponseOkMessage},
  }
  TestClient.DELETE(fmt.Sprintf("%s/deleteAdmin/%d", url, id), pageKeys)
}

```

## 🔋 JetBrains OS licenses

<a href="https://www.jetbrains.com/?from=iris-admin" target="_blank"><img src="https://raw.githubusercontent.com/panjf2000/illustrations/master/jetbrains/jetbrains-variant-4.png" width="230" align="middle"/></a>

## ☕️ Buy me a coffee

> Please be sure to leave your name, GitHub account or other social media accounts when you donate by the following means so that I can add it to the list of donors as a token of my appreciation.
- [为爱发电](https://afdian.net/@snowlyg/plan)
- [donating](https://paypal.me/snowlyg?country.x=C2&locale.x=zh_XC)



