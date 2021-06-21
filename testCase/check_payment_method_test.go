package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIcheck_payment_method(t *testing.T){
	var apiName = "check_payment_method"
	handler.DoTest(t, apiName)
}

//resp是多层json，expectJson不明确
// boolean 怎么搞
