/*
Package container provides functionality to interact with containers.
*/
package container

import "io"

type BuildOptions struct {
	Dockerfile   string
	OutputStream io.Writer
}

type PullOptions struct {
	OutputStream io.Writer
}

type CreateOptions struct {
	Binds         []string
	HasCommand    bool
	Shell         string
	NumContainers int
}

type StartOptions struct {
	Commands []string
}

type LogsOptions struct {
	Stdout io.Writer
	Stderr io.Writer
}

type WaitOptions struct{}

type KillOptions struct{}

type RemoveOptions struct{}

type Client interface {
	Build(imageName string, contextDir string, options BuildOptions) error
	Pull(imageName string, options PullOptions) error
	Create(imageName string, options CreateOptions) ([]string, error)
	Start(containerID string, options StartOptions) error
	Logs(containerID string, options LogsOptions) error
	Wait(containerID string, options WaitOptions) error
	Kill(containerID string, options KillOptions) error
	Remove(containerID string, options RemoveOptions) error
}

type DockerClientOptions struct {
	Host             string
	DockerTLSOptions *DockerTLSOptions
}

type DockerTLSOptions struct {
	CertPEMBlock []byte
	KeyPEMBlock  []byte
	CaPEMCert    []byte
}

func NewDockerClient(dockerClientOptions DockerClientOptions) (Client, error) {
	return newDockerClient(dockerClientOptions)
}
