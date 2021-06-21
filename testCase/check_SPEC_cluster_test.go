package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIcheck_SPEC_cluster(t *testing.T){
	var apiName = "check_SPEC_cluster"
	handler.DoTest(t, apiName)
}

//total int