package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIcheck_backups_by_cluster(t *testing.T){
	var apiName = "check_backups_by_cluster"
	handler.DoTest(t, apiName)
}

//expectJson.total 是 int