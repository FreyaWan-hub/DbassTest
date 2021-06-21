package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIcheck_SPEC_tidb(t *testing.T){
	var apiName = "check_SPEC_tidb"
	handler.DoTest(t, apiName)
}
//total int