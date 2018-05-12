package main

import (
	"fmt"
	"log"
	"net"

	"github.com/casalettoj/consensusprotocol/networking"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// private type for Context keys
type contextKey int

const (
	clientIDKey contextKey = iota
)

func authenticateClient(ctx context.Context, s *Server) (string, error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Printf("metadata: %v", md)
	}
	return "authenticated", nil
}

// unaryInterceptor calls authenticateClient with current context
// implements https://godoc.org/google.golang.org/grpc#UnaryInterceptor interface
func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	s, ok := info.Server.(*Server)
	if !ok {
		return nil, fmt.Errorf("type assertion failure for server")
	}
	clientID, err := authenticateClient(ctx, s)
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, clientIDKey, clientID)
	return handler(ctx, req)
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := Server{}

	creds, err := credentials.NewServerTLSFromFile("consensusprotocol.crt", "consensusprotocol.key")
	if err != nil {
		log.Fatalf("failed to generate credentials: %v", err)
	}

	opts := []grpc.ServerOption{grpc.Creds(creds), grpc.UnaryInterceptor(unaryInterceptor)}

	grpc := grpc.NewServer(opts...)
	networking.RegisterPingServer(grpc, &s)

	if err := grpc.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *networking.PingMessage) (*networking.PingMessage, error) {
	log.Printf("Receive message %s", in.Greeting)
	return &networking.PingMessage{Greeting: "bar"}, nil
}
