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

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

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
	router := router.GetRouter()
	return http.Serve(listener, router)
}
