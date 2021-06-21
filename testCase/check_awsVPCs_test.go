package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIcheck_awsVPCs(t *testing.T){
	var apiName = "check_awsVPCs"
	handler.DoTest(t, apiName)
}

//total应该是int