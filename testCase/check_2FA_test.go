package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIcheck_2FA(t *testing.T){
	var apiName = "check_2FA"
	handler.DoTest(t, apiName)
}
//path和para可以为空