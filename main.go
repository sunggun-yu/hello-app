package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sunggun-yu/hello-app/internal/config"
	"github.com/sunggun-yu/hello-app/internal/routers"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {
	// r := gin.Default()

	// Get port from environment variable or default to 8080
	port := config.Config.Port

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: routers.DefaultRouter(),
	}

	g.Go(func() error {
		return server.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
