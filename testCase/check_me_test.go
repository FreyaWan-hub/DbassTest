package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIcheck_me(t *testing.T){
	var apiName = "check_me"
	handler.DoTest(t, apiName)
}