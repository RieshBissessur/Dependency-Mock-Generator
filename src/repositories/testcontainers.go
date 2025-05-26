package repositories

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
)

func CreateTestContainer(image string, tag string, port string, name string) (string, error) {
	req := testcontainers.ContainerRequest{
		Name:         fmt.Sprintf("%s_%s", name, strconv.Itoa(rand.Intn(9999))),
		Image:        fmt.Sprintf("%s:%s", image, tag),
		ExposedPorts: []string{port},
		AutoRemove:   true,
	}

	container, containerErr := testcontainers.GenericContainer(context.Background(), testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if containerErr != nil {
		return "", containerErr
	}

	time.Sleep(time.Second)
	natPort, portError := nat.NewPort("tcp", port)
	if portError != nil {
		return "", portError
	}

	mappedPort, mappedPortError := container.MappedPort(context.Background(), natPort)
	if mappedPortError != nil {
		return "", mappedPortError
	}

	host, hostError := container.Host(context.Background())
	if hostError != nil {
		return "", hostError
	}

	return fmt.Sprintf("http://%s:%s", host, strconv.Itoa(mappedPort.Int())), nil
}

func CreateWireMockContainer(name string) (string, error) {
	return CreateTestContainer("wiremock/wiremock", "latest", "8080", name)
}
