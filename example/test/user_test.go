package test

import (
	"testing"
)

func TestList(t *testing.T) {
	client := TestServer.GetTestLogin(t, "/api/v1/auth/login", nil)
	defer client.Logout("/api/v1/users/logout", nil)

	// url := "v1/admin/api/getApiList"
	// pageKeys := tests.Responses{
	// 	{Key: "pageSize", Value: 10},
	// 	{Key: "page", Value: 1},
	// 	{Key: "list", Value: nil},
	// 	{Key: "total", Value: source.BaseApisLen()},
	// }
	// base.PostList(auth, url, base.PageRes, http.StatusOK, "获取成功", pageKeys)
}
