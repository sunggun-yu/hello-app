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

	// config for primary web server
	webConfig1 := config.WebConfig1()
	// config for secondary web server
	webConfig2 := config.WebConfig2()

	// kill application if port number are same
	if webConfig1.Port == webConfig2.Port {
		log.Fatal("Web port 1 and 2 are same")
	}

	// run primary web server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webConfig1.Port),
		Handler: routers.DefaultRouter(webConfig1),
	}
	g.Go(func() error {
		return server.ListenAndServe()
	})

	// run secondary web server if PORT_2 is specified
	if webConfig2.Port != "" {
		server2 := &http.Server{
			Addr:    fmt.Sprintf(":%s", webConfig2.Port),
			Handler: routers.DefaultRouter(webConfig2),
		}
		g.Go(func() error {
			return server2.ListenAndServe()
		})
	}

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
