package controllers

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func LogContainer(container string) string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	options := types.ContainerLogsOptions{ShowStdout: true}
	out, err := cli.ContainerLogs(ctx, container, options)
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, out)
	return os.Stdout
}
