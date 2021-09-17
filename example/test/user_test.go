package test

import (
	"testing"

	"github.com/snowlyg/helper/tests"
)

func TestList(t *testing.T) {
	client := TestServer.GetTestLogin(t, "/api/v1/auth/login", nil)
	defer client.Logout("/api/v1/users/logout", nil)
	url := "/api/v1/users"
	pageKeys := tests.Responses{
		{Key: "code", Value: 2000},
		{Key: "message", Value: "请求成功"},
		{Key: "data", Value: tests.Responses{
			{Key: "pageSize", Value: 10},
			{Key: "page", Value: 1},
			{Key: "items", Value: []tests.Responses{
				{
					{Key: "id", Value: 1, Type: "ge"},
					{Key: "name", Value: "超级管理员"},
					{Key: "username", Value: "admin"},
					{Key: "introduction", Value: "超级管理员"},
					{Key: "avatar", Value: "/static/images/avatar.jpg"},
					{Key: "roles", Value: []string{"超级管理员"}},
					{Key: "updatedAt", Value: "", Type: "notempty"},
					{Key: "createdAt", Value: "", Type: "notempty"},
				},
			}},
			{Key: "total", Value: 0, Type: "ge"},
		}},
	}
	client.GET(url, pageKeys, tests.RequestParams)
}

func TestCreate(t *testing.T) {
	client := TestServer.GetTestLogin(t, "/api/v1/auth/login", nil)
	defer client.Logout("/api/v1/users/logout", nil)
	url := "/api/v1/users"
	pageKeys := tests.Responses{
		{Key: "code", Value: 2000},
		{Key: "message", Value: "请求成功"},
		{Key: "data", Value: tests.Responses{
			{Key: "id", Value: 1, Type: "ge"},
		},
		},
	}
	data := map[string]interface{}{
		"name":     "测试名称",
		"username": "test_username",
		"intro":    "测试描述信息",
		"avatar":   "",
		"password": "123456",
	}

	client.POST(url, pageKeys, data)
}
