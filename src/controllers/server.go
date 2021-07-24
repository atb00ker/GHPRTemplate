package main

import (
	"ghpr/src/controllers/ghpr"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(ghpr.InsertGhprPath, ghpr.InsertGhprAction)
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
