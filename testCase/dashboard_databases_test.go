package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIdashboard_databases(t *testing.T){
	var apiName = "dashboard_databases"
	handler.DoTest(t, apiName)
}
//返回的不是json，csv怎么写不明确