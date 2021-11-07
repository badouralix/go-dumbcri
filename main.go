package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	// Using v1alpha2 instead of v1 to match the version used in crictl
	// See https://github.com/kubernetes-sigs/cri-tools/blob/36e98a6/cmd/crictl/version.go#L30
	// See https://github.com/kubernetes-sigs/cri-tools/pull/712
	cri "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

// https://github.com/kubernetes/kubernetes/blob/41ab2cc/staging/src/k8s.io/cri-api/pkg/apis/runtime/v1/api.proto
type RuntimeServiceServer struct {
	// Needed to compile even if some methods are not implemented here
	// See https://stackoverflow.com/questions/40823315/x-does-not-implement-y-method-has-a-pointer-receiver
	cri.UnimplementedRuntimeServiceServer
}

// Version serves `crictl version`
func (r *RuntimeServiceServer) Version(ctx context.Context, req *cri.VersionRequest) (*cri.VersionResponse, error) {
	log.Printf("Received request on Version with input %s", req.GetVersion())
	resp := cri.VersionResponse{
		Version:           req.GetVersion(),
		RuntimeName:       "go-dumbcri",
		RuntimeVersion:    "v0.0.0-alpha.1",
		RuntimeApiVersion: "v0.0.0-alpha.1",
	}
	return &resp, nil
}

type ImageServiceServer struct {
	// Needed to compile even if some methods are not implemented here
	// See https://stackoverflow.com/questions/40823315/x-does-not-implement-y-method-has-a-pointer-receiver
	cri.UnimplementedImageServiceServer
}

func main() {
	// See https://grpc.io/docs/languages/go/basics/
	server := grpc.NewServer()
	cri.RegisterRuntimeServiceServer(server, &RuntimeServiceServer{})
	cri.RegisterImageServiceServer(server, &ImageServiceServer{})

	// See https://stackoverflow.com/questions/2886719/unix-sockets-in-go
	lis, err := net.Listen("unix", "/tmp/go-dumbcri.sock")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	server.Serve(lis)
}
