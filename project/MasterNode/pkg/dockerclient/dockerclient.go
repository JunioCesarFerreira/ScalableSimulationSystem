package dockerclient

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

// NewDockerClient cria um novo cliente Docker
func NewDockerClient() (*client.Client, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("failed to create Docker client: %v", err)
	}
	return cli, nil
}

// CreateAndStartContainer cria e inicia um container Docker com nome e porta customizados
func CreateContainer(ctx context.Context, cli *client.Client, containerName string, hostPort string) (string, error) {
	containerConfig := &container.Config{
		Image: "ubuntu-docker",
		ExposedPorts: nat.PortSet{
			"22/tcp": struct{}{},
		},
	}

	networkName := "simnet"

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"22/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: hostPort,
				},
			},
		},
		NetworkMode: container.NetworkMode(networkName),
	}

	// Cria o container com nome personalizado
	containerResp, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, containerName)
	if err != nil {
		return "", fmt.Errorf("failed to create container: %v", err)
	}

	return containerResp.ID, nil
}

func StartContainer(ctx context.Context, cli *client.Client, containerID string) error {
	// Inicia o container
	err := cli.ContainerStart(ctx, containerID, container.StartOptions{})
	if err != nil {
		return fmt.Errorf("failed to start container: %v", err)
	}

	return nil
}

// RemoveContainer remove um container Docker
func RemoveContainer(ctx context.Context, cli *client.Client, containerID string) error {
	err := cli.ContainerRemove(ctx, containerID, container.RemoveOptions{Force: true})
	if err != nil {
		return fmt.Errorf("failed to remove container: %v", err)
	}
	return nil
}

// IsContainerRunning verifica se um container está em execução
func IsContainerRunning(ctx context.Context, cli *client.Client, containerID string) (bool, error) {
	containerJSON, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		return false, fmt.Errorf("failed to inspect container: %v", err)
	}
	return containerJSON.State.Running, nil
}
