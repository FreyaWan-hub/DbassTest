package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIpost_pwd_verify(t *testing.T){
	var apiName = "post_pwd_verify"
	handler.DoTest(t, apiName)
}