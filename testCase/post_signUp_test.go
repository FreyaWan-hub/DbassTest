package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIpost_signUp(t *testing.T){
	var apiName = "post_signUp"
	handler.DoTest(t, apiName)
}