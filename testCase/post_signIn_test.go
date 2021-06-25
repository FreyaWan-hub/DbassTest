package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIpost_signIn(t *testing.T){
	var apiName = "post_signIn"
	handler.DoTest(t, apiName)
}