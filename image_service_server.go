package main

import (
	"context"
	"log"
	"strconv"

	cri "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

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
