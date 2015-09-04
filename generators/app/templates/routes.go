package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Policy      func(http.Handler) http.Handler
}

type Routes []Route

var routes = Routes{
	Route{
		"Example Route Descriptive Title",
		"GET",
		"/",
		checkHealth,
		nil,
	},
}
