package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPItier_intro(t *testing.T){
	var apiName = "tier_intro"
	handler.DoTest(t, apiName)
}