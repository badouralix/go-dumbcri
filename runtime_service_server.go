package main

import (
	"context"
	"log"

	cri "k8s.io/cri-api/pkg/apis/runtime/v1"
)

type RuntimeServiceServer struct {
	// Needed to compile even if some methods are not implemented here
	// See https://stackoverflow.com/questions/40823315/x-does-not-implement-y-method-has-a-pointer-receiver
	cri.UnimplementedRuntimeServiceServer
}

// Version serves `crictl version`
func (r *RuntimeServiceServer) Version(ctx context.Context, req *cri.VersionRequest) (*cri.VersionResponse, error) {
	log.Printf("Received request on Version with input %s", req.GetVersion())
	return &cri.VersionResponse{
		Version:           req.GetVersion(),
		RuntimeName:       "go-dumbcri",
		RuntimeVersion:    "v0.0.0-alpha.1",
		RuntimeApiVersion: "v0.0.0-alpha.1",
	}, nil
}
