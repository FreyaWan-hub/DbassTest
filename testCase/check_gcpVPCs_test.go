package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIcheck_gcpVPCs(t *testing.T){
	var apiName = "check_gcpVPCs"
	handler.DoTest(t, apiName)
}
//total int
