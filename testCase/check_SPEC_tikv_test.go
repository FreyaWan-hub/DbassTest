package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIcheck_SPEC_tikv(t *testing.T){
	var apiName = "check_SPEC_tikv"
	handler.DoTest(t, apiName)
}
//total int