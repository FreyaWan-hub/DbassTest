package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIdashboard_check_heatmaps(t *testing.T){
	var apiName = "dashboard_check_heatmaps"
	handler.DoTest(t, apiName)
}

//参数未知