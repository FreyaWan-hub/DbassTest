package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIcheck_SPEC_tiflash(t *testing.T){
	var apiName = "check_SPEC_tiflash"
	handler.DoTest(t, apiName)
}

//total int
