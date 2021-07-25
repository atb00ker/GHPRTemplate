package main

import (
	"ghpr/src/controllers/ghpr"
)

var httpRouter Router = NewMuxRouter()

func main() {
	// GHPR Routes
	ghprCaller := ghpr.GhprCaller{Hasura: &ghpr.HasuraClient{}}
	httpRouter.POST(ghpr.InsertGhprPath, ghprCaller.InsertGhprAction)
	// Start Server
	httpRouter.SERVE(":3000")
}
