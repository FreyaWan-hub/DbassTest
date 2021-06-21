package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIdashboard_conf(t *testing.T){
	var apiName = "dashboard_info"
	handler.DoTest(t, apiName)
}
//boolean