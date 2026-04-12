package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	pb "github.com/sunggun-yu/hello-app/grpc"
	"github.com/sunggun-yu/hello-app/internal/config"
	"github.com/sunggun-yu/hello-app/internal/routers"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	// config for primary web server
	webConfig1 := config.WebConfig1()
	// config for secondary web server
	webConfig2 := config.WebConfig2()
	// config for primary grpc server
	grpcConfig1 := config.GrpcConfig1()

	// kill application if port numbers conflict
	if webConfig1.Port == webConfig2.Port || webConfig1.Port == grpcConfig1.Port || webConfig2.Port == grpcConfig1.Port {
		log.Fatal("Port conflict detected between servers")
	}

	// context that cancels on SIGINT/SIGTERM for graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	g, ctx := errgroup.WithContext(ctx)

	// run primary web server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webConfig1.Port),
		Handler: routers.DefaultRouter(webConfig1),
	}
	g.Go(func() error {
		log.Printf("starting primary web server on :%s", webConfig1.Port)
		return server.ListenAndServe()
	})

	// run secondary web server
	server2 := &http.Server{
		Addr:    fmt.Sprintf(":%s", webConfig2.Port),
		Handler: routers.DefaultRouter(webConfig2),
	}
	g.Go(func() error {
		log.Printf("starting secondary web server on :%s", webConfig2.Port)
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
		log.Printf("starting grpc server on :%s", grpcConfig1.Port)
		return grpcServer.Serve(lis)
	})

	// graceful shutdown goroutine
	g.Go(func() error {
		<-ctx.Done()
		log.Println("shutting down servers gracefully...")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		grpcServer.GracefulStop()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("primary web server shutdown error: %v", err)
		}
		if err := server2.Shutdown(shutdownCtx); err != nil {
			log.Printf("secondary web server shutdown error: %v", err)
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}
	log.Println("all servers stopped")
}
