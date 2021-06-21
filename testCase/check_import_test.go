package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIcheck_import(t *testing.T){
	var apiName = "check_import"
	handler.DoTest(t, apiName)
}
// totalfile int