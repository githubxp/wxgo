package main

import (
	"log"
	"net/http"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()

		inner.ServeHTTP(rw, req)

		var ip = req.Header.Get("Remote_addr")
		if ip == "" {
			ip = req.RemoteAddr
		}

		log.Printf(
			"%s\t%s\t%s\t%s\t%s",
			ip,
			req.Method,
			req.RequestURI,
			name,
			time.Since(start),
		)
	})
}
