package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIcheck_RB_by_orgId(t *testing.T){
	var apiName = "check_RB_by_orgId"
	handler.DoTest(t, apiName)
}
//int total