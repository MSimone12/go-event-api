package main

import (
	"fmt"
	"go-event-api/src/router"
	"log"
	"net/http"
	"os"
)

func main() {
	router := router.GetRouter()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))
}
