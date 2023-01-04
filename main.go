package main

import (
	"Calculator/api"
	"log"
	"net/http"
)

func main() {
	api.InitializeEndpoints()
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
