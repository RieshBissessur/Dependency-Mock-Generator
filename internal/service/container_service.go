package service

import (
	"fmt"

	"github.com/rieshbissessur/dependency-mock-generator/internal/model"
	"github.com/rieshbissessur/dependency-mock-generator/internal/repository"
)

func SetupContainer(container model.Container) (string, error) {
	url, err := repository.CreateTestContainer(container.Image, container.Tag, container.Port, container.Name)
	if err != nil {
		fmt.Println("Error creating WireMock test container:", err)
		return "", err
	}

	fmt.Printf("%s WireMock container is available on %s/__admin/mappings", container.Name, url)
	return url, nil
}
