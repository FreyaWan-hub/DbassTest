
package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIpost_auto_backup(t *testing.T){
	var apiName = "post_auto_backup"
	handler.DoTest(t, apiName)
}