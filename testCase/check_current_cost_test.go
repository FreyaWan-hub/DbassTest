package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIcheck_current_cost(t *testing.T){
	var apiName = "check_current_cost"
	handler.DoTest(t, apiName)
}
//expectJson不明确
