package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIcheck_user_access(t *testing.T){
	var apiName = "check_user_access"
	handler.DoTest(t, apiName)
}