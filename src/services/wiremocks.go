package services

import (
	"encoding/json"
	"fmt"
	"mock-generator/models"
	"mock-generator/providers"
	"mock-generator/repositories"
	"time"
)

func SetupMock(wireMocks providers.Mock) (string, error) {
	url, err := repositories.CreateWireMockContainer(wireMocks.Name)

	if err != nil {
		fmt.Println("Error creating WireMock test container:", err)
		return "", err
	}

	return url, nil
}

func ExportMockState(name string, url string) error {
	mappingsString, err := repositories.GetAllMappingModels(url)
	if err != nil {
		return err
	}

	var mappings models.Mappings
	if err := json.Unmarshal([]byte(mappingsString), &mappings); err != nil {
		return err
	}

	mappings.Name = &name

	fileName := fmt.Sprintf("%s-%s.json", name, time.Now().Format("2006-01-02T15:04:05Z07:00"))
	err = repositories.WriteFileContentFromAny("mappings", fileName, mappings)
	return nil
}

func ImportMockState(pathToFile string) (string, string, error) {
	content, readError := repositories.ReadFileContentAsBytes(pathToFile)
	if readError != nil {
		return "", "", readError
	}

	var mappings models.Mappings
	if err := json.Unmarshal(content, &mappings); err != nil {
		return "", "", err
	}

	mock := providers.Mock{Name: *mappings.Name}
	mockUrl, err := SetupMock(mock)
	if err != nil {
		return "", "", err
	}

	importError := repositories.ImportMappingModels(mockUrl, mappings.Mappings[0])
	if importError != nil {
		return "", "", importError
	}

	return mockUrl, *mappings.Name, nil
}
