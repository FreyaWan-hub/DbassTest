package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIdashboard_time_ranges(t *testing.T){
	var apiName = "dashboard_time_ranges"
	handler.DoTest(t, apiName)
}