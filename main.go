package main

import (
	"log"
	"net"
	"syscall"

	"google.golang.org/grpc"

	cri "k8s.io/cri-api/pkg/apis/runtime/v1"
)

const (
	socketaddr = "/tmp/go-dumbcri.sock"
)

func main() {
	// See https://grpc.io/docs/languages/go/basics/
	server := grpc.NewServer()
	cri.RegisterRuntimeServiceServer(server, &RuntimeServiceServer{})
	cri.RegisterImageServiceServer(server, &ImageServiceServer{Images: make(map[string]*cri.Image)})

	// We might want to acquire a lock instead of unlinking
	// See https://gavv.github.io/articles/unix-socket-reuse/
	err := syscall.Unlink(socketaddr)
	if err != nil {
		log.Printf("Failed to unlink: %v", err)
	}
	// See https://stackoverflow.com/questions/2886719/unix-sockets-in-go
	lis, err := net.Listen("unix", socketaddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Print(server.Serve(lis))
}
