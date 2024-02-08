package controllers

import (
	"bytes"
	"context"
	"fmt"
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

	var buf bytes.Buffer
	fmt.Fprint(&buf, out)
	os.Stdout = os.NewFile(0, "/dev/stdout")
	return buf.String()
}
