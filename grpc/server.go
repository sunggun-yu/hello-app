package grpc

import (
	context "context"
	"log"

	"github.com/sunggun-yu/hello-app/internal/config"
	helloService "github.com/sunggun-yu/hello-app/internal/service"
)

var cfg = config.GrpcConfig1()

type helloServiceServer struct {
	UnimplementedHelloServiceServer
}

func (s *helloServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	log.Printf("[GRPC]\tReceived Request - Ping")
	data := helloService.Ping(cfg)
	resp := PingResponse{
		Message:   data.Message,
		Service:   data.Service,
		Version:   data.Version,
		Instance:  data.Instance,
		Timestamp: data.Timestamp,
	}
	return &resp, nil
}

// SayHello implements
func (s *helloServiceServer) SayHello(_ context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("[GRPC]\tReceived Request - SayHello: %v", in.GetName())
	return &HelloReply{Message: "Hello " + in.GetName()}, nil
}

func NewHelloServiceServer() HelloServiceServer {
	return &helloServiceServer{}
}
