package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIdashboard_statement_list(t *testing.T){
	var apiName = "dashboard_statement_list"
	handler.DoTest(t, apiName)
}
//返回参数不明确