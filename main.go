package main

import (
	"fmt"
	"go-event-api/router"
	"log"
	"net/http"
	"os"
)

func main() {
	router := router.GetRouter()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), router))
}
