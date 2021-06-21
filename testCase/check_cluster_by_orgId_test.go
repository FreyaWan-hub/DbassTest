package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIcheck_cluster_by_orgId(t *testing.T){
	var apiName = "check_cluster_by_orgId"
	handler.DoTest(t, apiName)
}
// int类型