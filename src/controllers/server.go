package main

import (
	"ghpr/src/controllers/ghpr"
	"net/http"
)

var httpRouter Router = NewMuxRouter()

func main() {
	// GHPR Routes
	ghprCaller := ghpr.GhprCaller{Hasura: &ghpr.HasuraClient{}}
	httpRouter.POST(ghpr.InsertGhprPath, ghprCaller.InsertGhprAction)
	// Static Files
	muxDispatcher.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist/")))
	// Start Server
	httpRouter.SERVE(":3000")
}
