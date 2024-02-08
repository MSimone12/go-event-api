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

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), router))
}