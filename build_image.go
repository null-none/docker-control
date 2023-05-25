package main

import (
	"log"
	"io/ioutil"
	"io"
	"context"
	"bytes"
	"os"
	"archive/tar"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func buildImage(client *client.Client, tags []string, dockerfile string)  error {
	ctx := context.Background()

	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	dockerFileReader, err := os.Open(dockerfile)
	if err != nil {
		return err
	}

	readDockerFile, err := ioutil.ReadAll(dockerFileReader)
	if err != nil {
		return err
	}

	tarHeader := &tar.Header{
		Name: dockerfile,
		Size: int64(len(readDockerFile)),
	}

	err = tw.WriteHeader(tarHeader)
    if err != nil {
		return err
    }

    _, err = tw.Write(readDockerFile)
    if err != nil {
		return err
    }

    dockerFileTarReader := bytes.NewReader(buf.Bytes())

	buildOptions := types.ImageBuildOptions{
        Context:    dockerFileTarReader,
        Dockerfile: dockerfile,
        Remove:     true,
		Tags: 		tags,
	}


	imageBuildResponse, err := client.ImageBuild(
        ctx,
        dockerFileTarReader,
		buildOptions,
	)

	if err != nil {
		return err
	}

	defer imageBuildResponse.Body.Close()
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	if err != nil {
		return err
	}

	return nil
}
