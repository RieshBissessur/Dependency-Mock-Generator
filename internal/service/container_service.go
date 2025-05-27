package service

import (
	"fmt"

	"github.com/rieshbissessur/dependency-mock-generator/internal/model"
	"github.com/rieshbissessur/dependency-mock-generator/internal/repository"
)

func SetupContainer(container model.Container) (string, error) {
	url, err := repository.CreateTestContainer(container.Image, container.Tag, container.Port, container.Name)
	if err != nil {
		fmt.Printf("Error creating test container for %s: %s", container.Name, err)
		return "", err
	}

	fmt.Printf("%s container is available on %s/__admin/mappings", container.Name, url)
	return url, nil
}
