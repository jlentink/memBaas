package main

import (
	"context"
	log "github.com/jlentink/yaglogger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"memBaas/internal/cache"
	"memBaas/internal/cnf"
	"memBaas/internal/proto"
	"net"
	"time"
)

type server struct {
	proto.UnsafeServerServer
}

func (s *server) Get(_ context.Context, in *proto.GetRequest) (*proto.GetResponse, error) {
	get, status := cache.Get(in.Key)

	return &proto.GetResponse{
		Status: status.GetStatus(),
		Value:  get,
	}, nil
}

func (s *server) Set(_ context.Context, in *proto.SetRequest) (*proto.SetResponse, error) {
	set := cache.Set(in.Key, in.Value, in.Ttl, in.Overwrite)
	return &proto.SetResponse{
		Status:  set.SetStatus(),
		Message: "",
	}, nil
}

func (s *server) Delete(_ context.Context, in *proto.GetRequest) (*proto.GetResponse, error) {
	del := cache.Delete(in.Key)
	return &proto.GetResponse{
		Status: del.GetStatus(),
	}, nil
}

func main() {
	log.SetLevel(log.LevelAll)
	cnf.Get()
	log.SetLevelByString(cnf.GetString("log.level"))
	go garbageRun()

	sock, err := net.Listen("tcp", cnf.GetString("network.address"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	proto.RegisterServerServer(s, &server{})
	log.Printf("memBaas listening at %v", sock.Addr())
	if err := s.Serve(sock); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func garbageRun() {
	for {
		cache.Cleanup()
		time.Sleep(1 * time.Second)
	}
}
