package main

import (
	"github.com/tidbcloud/qa-Gourd/handler"
	"testing"
)

func TestAPIdashboard_check_heatmaps_config(t *testing.T){
	var apiName = "dashboard_check_heatmaps_config"
	handler.DoTest(t, apiName)
}

//expectJson.auto_collection_disabled boolean