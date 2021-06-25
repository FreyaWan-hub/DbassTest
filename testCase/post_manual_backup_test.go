package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIpost_manual_backup(t *testing.T){
	var apiName = "post_manual_backup"
	handler.DoTest(t, apiName)
}