package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIcheck_traffic_filter(t *testing.T){
	var apiName = "check_traffic_filter"
	handler.DoTest(t, apiName)
}
//total int