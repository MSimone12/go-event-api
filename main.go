package main

import (
	"go-event-api/router"
	"log"
	"net/http"
)

func main() {
	router := router.GetRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
