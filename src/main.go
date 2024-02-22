package main

import (
	"context"
	"fmt"
	"go-event-api/src/router"
	"log"
	"net/http"

	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

// This is the entrypoint
func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

// run will run ngrok proxy serving with router
func run(ctx context.Context) error {
	listener, err := ngrok.Listen(
		ctx,
		config.HTTPEndpoint(config.WithDomain("pleasing-ample-lab.ngrok-free.app")),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		return err
	}
	fmt.Println("App URL", listener.URL())
	// router contains the all the route logic,
	// which will redirect to the right handlers
	router := router.GetRouter()
	return http.Serve(listener, router)
}
