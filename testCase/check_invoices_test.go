package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIcheck_invoices(t *testing.T){
	var apiName = "check_invoices"
	handler.DoTest(t, apiName)
}

//expectJson不明确
