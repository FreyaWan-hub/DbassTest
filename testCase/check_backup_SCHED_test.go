package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIcheck_backup_SCHED(t *testing.T){
	var apiName = "check_backup_SCHED"
	handler.DoTest(t, apiName)
}