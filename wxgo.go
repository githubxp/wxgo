package main

import (
	"net/http"
	"log"
)

func main() {
	router := NewRouter()
	err := http.ListenAndServe(":80", router)

	if err != nil {
		log.Fatal(err)
	}
}
