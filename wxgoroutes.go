package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"SHOW_HELLO",
		"GET",
		"/",
		ShowHello,
	},
	Route{
		"JSSDK_API",
		"GET",
		"/appid/{appid}",
		Jssdk_api,
	},
}
