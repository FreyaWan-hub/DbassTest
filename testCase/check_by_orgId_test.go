package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIcheck_by_orgId(t *testing.T){
	var apiName = "check_by_orgId"
	handler.DoTest(t, apiName)
}
//expectJson.is_org_owner 需要boolean格式