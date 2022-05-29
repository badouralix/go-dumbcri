package main

import (
	"context"
	"encoding/hex"
	"errors"
	"log"
	"strconv"
	"strings"

	"golang.org/x/crypto/sha3"

	cri "k8s.io/cri-api/pkg/apis/runtime/v1"
)

type ImageServiceServer struct {
	// Images is a map of image id to cri images.
	Images map[string]*cri.Image
}

// ListImages serves `crictl images`
func (i *ImageServiceServer) ListImages(ctx context.Context, req *cri.ListImagesRequest) (*cri.ListImagesResponse, error) {
	log.Printf("Received request on ListImages with Filter=%s", req.GetFilter())

	images := make([]*cri.Image, 0)
	for _, image := range i.Images {
		images = append(images, image)
	}

	return &cri.ListImagesResponse{
		Images: images,
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

	split := strings.Split(req.Image.Image, ":")
	if len(split) > 2 {
		return nil, errors.New("request did not match expected image:tag format")
	}

	imageName := split[0]
	imageTag := "latest"
	if len(split) == 2 {
		imageTag = split[1]
	}
	image := imageName + ":" + imageTag

	h := sha3.New512()
	h.Write([]byte(image))
	id := hex.EncodeToString(h.Sum(nil))

	i.Images[imageName] = &cri.Image{
		Id:          id,
		RepoTags:    []string{image},
		RepoDigests: []string{},
		Size_:       0,
		Spec:        req.Image,
	}

	return &cri.PullImageResponse{
		ImageRef: id,
	}, nil
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
