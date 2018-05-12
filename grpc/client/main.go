package main

import (
	"log"

	"github.com/casalettoj/consensusprotocol/auth"
	"github.com/casalettoj/consensusprotocol/networking"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	var conn *grpc.ClientConn

	creds, err := credentials.NewClientTLSFromFile("consensusprotocol.crt", "")
	if err != nil {
		log.Fatalf("could not load ssl cert: %s", err)
	}

	auth := auth.Auth{
		User: "Username",
	}

	conn, err = grpc.Dial("localhost:7777", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := networking.NewPingClient(conn)
	response, err := c.SayHello(context.Background(), &networking.PingMessage{Greeting: "foo"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from server: %s", response.Greeting)
}
