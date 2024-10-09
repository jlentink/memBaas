package main

import (
	"context"
	log "github.com/jlentink/yaglogger"
	"google.golang.org/grpc"
	"memBaas/internal/proto"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := proto.NewServerClient(conn)

	resp, err := client.Set(context.Background(), &proto.SetRequest{
		Key:       "key",
		Value:     "value",
		Ttl:       10,
		Overwrite: true,
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Info("Response: %s", resp)

	getResp, err := client.Get(context.Background(), &proto.GetRequest{Key: "key"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Info("Response: %s", getResp.Value)

	client.Set(context.Background(), &proto.SetRequest{
		Key:       "bla",
		Value:     "bla",
		Ttl:       5,
		Overwrite: false,
	})

	client.Delete(context.Background(), &proto.GetRequest{Key: "bla"})
}
