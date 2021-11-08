package main

import (
	"context"
	"log"
	"net"
	"strconv"

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
	return &cri.VersionResponse{
		Version:           req.GetVersion(),
		RuntimeName:       "go-dumbcri",
		RuntimeVersion:    "v0.0.0-alpha.1",
		RuntimeApiVersion: "v0.0.0-alpha.1",
	}, nil
}

type ImageServiceServer struct {
	Images []*cri.Image
}

// ListImages serves `crictl images`
func (i *ImageServiceServer) ListImages(ctx context.Context, req *cri.ListImagesRequest) (*cri.ListImagesResponse, error) {
	log.Printf("Received request on ListImages with Filter=%s", req.GetFilter())
	return &cri.ListImagesResponse{
		Images: i.Images,
	}, nil
}

// ImageStatus serves `crictl inspecti`
func (i *ImageServiceServer) ImageStatus(ctx context.Context, req *cri.ImageStatusRequest) (*cri.ImageStatusResponse, error) {
	log.Printf("Received request on ImageStatus with Image=%s Verbose=%s", req.GetImage(), strconv.FormatBool(req.GetVerbose()))
	return &cri.ImageStatusResponse{
		Image: &cri.Image{},
		Info: map[string]string{
			"key": "value",
		},
	}, nil
}

// PullImage serves `crictl pull`
func (i *ImageServiceServer) PullImage(ctx context.Context, req *cri.PullImageRequest) (*cri.PullImageResponse, error) {
	log.Printf("Received request on PullImage with Image=%s Auth=%s SandboxConfig=%s", req.GetImage(), req.GetAuth(), req.GetSandboxConfig())
	return &cri.PullImageResponse{}, nil
}

// RemoveImage serves `crictl rmi`
func (i *ImageServiceServer) RemoveImage(ctx context.Context, req *cri.RemoveImageRequest) (*cri.RemoveImageResponse, error) {
	log.Printf("Received request on RemoveImage with Image=%s", req.GetImage())
	return &cri.RemoveImageResponse{}, nil
}

// ImageFsInto serves `crictl imagefsinfo`
func (i *ImageServiceServer) ImageFsInfo(ctx context.Context, req *cri.ImageFsInfoRequest) (*cri.ImageFsInfoResponse, error) {
	log.Printf("Received request on ImageFsInfo")
	return &cri.ImageFsInfoResponse{
		ImageFilesystems: []*cri.FilesystemUsage{},
	}, nil
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
