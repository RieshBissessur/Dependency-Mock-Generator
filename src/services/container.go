package services

import (
	"fmt"
	"mock-generator/providers"
	"mock-generator/repositories"
)

func SetupContainer(container providers.Container) (string, error) {
	url, err := repositories.CreateTestContainer(container.Image, container.Tag, container.Port, container.Name)
	if err != nil {
		fmt.Println("Error creating WireMock test container:", err)
		return "", err
	}

	fmt.Printf("%s WireMock container is available on %s/__admin/mappings", container.Name, url)
	return url, nil
}
