package main

import (
	"log"
	"net/http"
)

var app AppConfig
var th ThoughtHandler

func main() {
	th = newThoughtHandler()
	log.Println("running on port 9090")
	log.Fatal(http.ListenAndServe(":9090", http.HandlerFunc(Serve)))
}
