package main

import (
	"encoding/json"
	"net/http"
)

type jsonErr struct {
	Code bool   `json:"auth"`
	Text string `json:"error"`
}

type jsonStatus struct {
	Status string `json:"status"`
}

func sendObject(w http.ResponseWriter, s int, o interface{}) {
	js, _ := json.Marshal(o)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)
	w.Write(js)
}

func sendM(w http.ResponseWriter, s int, a bool, m string) {
	js, _ := json.Marshal(jsonErr{a, m})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)
	w.Write(js)
}
