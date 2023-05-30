package controllers

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func PullImage(image string) string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	out, err := cli.ImagePull(ctx, image, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}

	defer out.Close()

	io.Copy(os.Stdout, out)
	return "success"
}
