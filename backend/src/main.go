package main

import (
	"net/http"
	"os"
	"time"
	"encoding/json"
)

type HandsOn struct {
	Time 	 time.Time `json:"time"`
	Hostname string    `json:"hostname`
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
	if r.URL.Path != "/" {
		w.Write([]byte("NO found"))
		return
	}
	resp := HandsOn{
		Time: time.Now(),
		Hostname: os.Getenv("HOSTNAME"),
	}

	jsonResp, err := json.Marshal(&resp)

	if err != nil {
		w.Write([]byte("Error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func main() {
	http.HandleFunc("/", ServeHTTP)
	http.ListenAndServe(":9090", nil)
}