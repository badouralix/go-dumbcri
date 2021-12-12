package main

import (
	"log"
	"net"
	"syscall"

	"google.golang.org/grpc"

	// Using v1alpha2 instead of v1 to match the version used in crictl
	// See https://github.com/kubernetes-sigs/cri-tools/blob/36e98a6/cmd/crictl/version.go#L30
	// See https://github.com/kubernetes-sigs/cri-tools/pull/712
	cri "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
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
