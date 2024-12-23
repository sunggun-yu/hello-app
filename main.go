package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/sunggun-yu/hello-app/grpc"
	"github.com/sunggun-yu/hello-app/internal/config"
	"github.com/sunggun-yu/hello-app/internal/routers"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	g errgroup.Group
)

func main() {

	// config for primary web server
	webConfig1 := config.WebConfig1()
	// config for secondary web server
	webConfig2 := config.WebConfig2()
	// config for secondary web server
	grpcConfig1 := config.GrpcConfig1()

	// kill application if port numbers conflict
	if webConfig1.Port == webConfig2.Port || webConfig1.Port == grpcConfig1.Port || webConfig2.Port == grpcConfig1.Port {
		log.Fatal("Port conflict detected between servers")
	}

	// run primary web server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webConfig1.Port),
		Handler: routers.DefaultRouter(webConfig1),
	}
	g.Go(func() error {
		return server.ListenAndServe()
	})

	// run secondary web server
	server2 := &http.Server{
		Addr:    fmt.Sprintf(":%s", webConfig2.Port),
		Handler: routers.DefaultRouter(webConfig2),
	}
	g.Go(func() error {
		return server2.ListenAndServe()
	})

	// run grpc server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcConfig1.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	helloGrpcServer := pb.NewHelloServiceServer()
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterHelloServiceServer(grpcServer, helloGrpcServer)
	// register reflection service on gRPC server.
	reflection.Register(grpcServer)

	g.Go(func() error {
		return grpcServer.Serve(lis)
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
