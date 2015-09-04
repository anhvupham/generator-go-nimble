package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheckHealth(t *testing.T) {
	r, _ := http.NewRequest("GET", "http://localhost:<%= port %>/", nil)
	w := httptest.NewRecorder()
	checkHealth(w, r)
	var js jsonStatus
	_ = json.Unmarshal([]byte(w.Body.String()), &js)
	if js.Status != "ok" {
		t.Error("check health failed")
		t.Fail()
	}
	if w.Code != 200 {
		t.Error("expected http code 200")
		t.Fail()
	}

	pass := false
	for _, v := range w.HeaderMap["Content-Type"] {
		if v == "application/json" {
			pass = true
		}
	}
	if pass != true {
		t.Error("expected json content type")
		t.Fail()
	}
}
