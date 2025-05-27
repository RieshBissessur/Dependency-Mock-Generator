package service

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/rieshbissessur/dependency-mock-generator/internal/model"
	"github.com/rieshbissessur/dependency-mock-generator/internal/repository"
)

func SetupMock(mock model.Mock) (string, error) {
	url, err := repository.CreateWireMockContainer(mock.Name)

	if err != nil {
		fmt.Println("Error creating WireMock test container:", err)
		return "", err
	}

	if mock.File == nil || *mock.File == "" {
		return url, nil
	}

	mappings, readError := ReadMapingModelFromFile(*mock.File)
	if readError != nil {
		return "", readError
	}

	importError := ImportMockState(url, mappings)
	if importError != nil {
		return "", importError
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

func CreateAndImportMockState(pathToFile string) (string, string, error) {
	mappings, readError := ReadMapingModelFromFile(pathToFile)
	if readError != nil {
		return "", "", readError
	}

	mock := model.Mock{Name: *mappings.Name}
	mockUrl, err := SetupMock(mock)
	if err != nil {
		return "", "", err
	}

	importError := ImportMockState(mockUrl, mappings)
	if importError != nil {
		return "", "", importError
	}

	return mockUrl, *mappings.Name, nil
}

func ImportMockState(mockUrl string, mappings model.Mappings) error {
	importError := repository.ImportMappingModels(mockUrl, mappings.Mappings[0])
	if importError != nil {
		return importError
	}

	return nil
}

func ReadMapingModelFromFile(pathToFile string) (model.Mappings, error) {
	content, readError := repository.ReadFileContentAsBytes(pathToFile)
	if readError != nil {
		return model.Mappings{}, readError
	}

	var mappings model.Mappings
	if serializationError := json.Unmarshal(content, &mappings); serializationError != nil {
		return model.Mappings{}, serializationError
	}

	return mappings, nil
}
