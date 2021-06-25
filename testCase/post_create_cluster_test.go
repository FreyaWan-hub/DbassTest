package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIpost_create_cluster(t *testing.T){
	var apiName = "post_create_cluster"

	handler.DoTest(t, apiName)
}