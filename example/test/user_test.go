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
		{Key: "pageSize", Value: 10},
		{Key: "page", Value: 1},
		{Key: "list", Value: nil},
		{Key: "total", Value: 0, Type: "GE"},
	}
	client.GET(url, pageKeys, tests.RequestParams)
}
