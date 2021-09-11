package test

import (
	"testing"
)

func TestList(t *testing.T) {
	TestServer.GetTestLogin(t, "/api/v1/login", nil)
	defer TestServer.GetTestLogout(t, "/api/v1/logout", nil)

	// url := "v1/admin/api/getApiList"
	// pageKeys := base.ResponseKeys{    
	// 	{Key: "pageSize", Value: 10},
	// 	{Key: "page", Value: 1},
	// 	{Key: "list", Value: nil},
	// 	{Key: "total", Value: source.BaseApisLen()},
	// }
	// base.PostList(auth, url, base.PageRes, http.StatusOK, "获取成功", pageKeys)
}
