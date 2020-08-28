package handler

import (
	"encoding/json"
	"github.com/marzunas/go-app/store"
	"testing"
)


func Test_getStatJSON(t *testing.T)  {
	store.ExecTimes = append(store.ExecTimes, 121121)
	var stats = getStatJSON()
	if stats == "" {
		t.Error("stats should be available")
	}
	var statResponse StatResponse
	var err = json.Unmarshal([]byte(stats), &statResponse)
	if err != nil || &statResponse == nil {
		t.Error("should have been valid json")
	}
}
