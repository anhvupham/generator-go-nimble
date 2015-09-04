package main

import "net/http"

func checkHealth(w http.ResponseWriter, r *http.Request) {
	sendObject(w, 200, jsonStatus{"ok"})
}
