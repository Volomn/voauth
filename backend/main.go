package main

import (
	"log"
	"net/http"

	"github.com/Volomn/voauth/backend/api"
)

func main() {
	// create api router
	router := api.GetApiRouter()

	// start server
	log.Fatal(http.ListenAndServe(":5000", router))
}
