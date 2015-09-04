package main

import "net/http"

func IsAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := rdb.Get(r, "session-key")
		if session.Values["loggedin"] == true {
			h.ServeHTTP(w, r)
		} else {
			sendM(w, 401, false, "User is not logged in")
		}
	})
}
