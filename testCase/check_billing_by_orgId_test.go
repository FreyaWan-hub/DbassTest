package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIcheck_billing_by_orgId(t *testing.T){
	var apiName = "check_billing_by_orgId"
	handler.DoTest(t, apiName)
}