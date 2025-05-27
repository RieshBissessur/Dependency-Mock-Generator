package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/rieshbissessur/dependency-mock-generator/internal/model"
	"github.com/rieshbissessur/dependency-mock-generator/internal/repository"
)

func SetupMock(wireMocks model.Mock) (string, error) {
	url, err := repository.CreateWireMockContainer(wireMocks.Name)

	if err != nil {
		fmt.Println("Error creating WireMock test container:", err)
		return "", err
	}

	return url, nil
}

func ExportMockState(name string, url string) error {
	mappingsString, err := repository.GetAllMappingModels(url)
	if err != nil {
		return err
	}

	var mappings model.Mappings
	if err := json.Unmarshal([]byte(mappingsString), &mappings); err != nil {
		return err
	}

	mappings.Name = &name

	fileName := fmt.Sprintf("%s-%s.json", name, time.Now().Format("2006-01-02T15:04:05Z07:00"))
	err = repository.WriteFileContentFromAny("mappings", fileName, mappings)
	return nil
}

func ImportMockState(pathToFile string) (string, string, error) {
	content, readError := repository.ReadFileContentAsBytes(pathToFile)
	if readError != nil {
		return "", "", readError
	}

	var mappings model.Mappings
	if err := json.Unmarshal(content, &mappings); err != nil {
		return "", "", err
	}

	mock := model.Mock{Name: *mappings.Name}
	mockUrl, err := SetupMock(mock)
	if err != nil {
		return "", "", err
	}

	importError := repository.ImportMappingModels(mockUrl, mappings.Mappings[0])
	if importError != nil {
		return "", "", importError
	}

	return mockUrl, *mappings.Name, nil
}
