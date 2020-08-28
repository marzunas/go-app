package handler

import (
	"encoding/json"
	"github.com/marzunas/go-app/store"
	"net/http"
)

type StatResponse struct {
	Total   int     `json:"total"`
	Average float64 `json:"average"`
}

func StatsHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(getStatJSON()))
		w.Write([]byte("\n"))
	})
}

func getStatJSON() string {
	var execTimes = store.ExecTimes
	var requestCount = len(execTimes)
	var totalTime int64 = 0
	for _, value := range execTimes {
		totalTime += value
	}
	var average = float64(totalTime) / float64(requestCount)
	response := &StatResponse{
		Total:   requestCount,
		Average: average,
	}
	result, _ := json.Marshal(response)
	return string(result)
}
