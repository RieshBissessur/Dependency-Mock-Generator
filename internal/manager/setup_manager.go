package manager

import (
	"github.com/rieshbissessur/dependency-mock-generator/internal/provider"
	"github.com/rieshbissessur/dependency-mock-generator/internal/service"
)

var ActiveContainers = make(map[string]string)
var ActiveMocks = make(map[string]string)

func RunSetup(setupPathFile string) error {

	setup, setupErr := provider.GetSetupFromFile(setupPathFile)
	if setupErr != nil {
		return nil
	}

	for _, container := range setup.Containers {
		containerUrl, err := service.SetupContainer(container)
		if err != nil {
			continue
		}

		ActiveContainers[container.Name] = containerUrl
	}

	for _, wireMock := range setup.Mocks {
		mockUrl, err := service.SetupMock(wireMock)
		if err != nil {
			continue
		}

		ActiveMocks[wireMock.Name] = mockUrl
	}

	return nil
}

func ExportMockStates() error {
	for name, wireMock := range ActiveMocks {
		exportError := service.ExportMockState(name, wireMock)
		if exportError != nil {
			return exportError
		}
	}

	return nil
}

func ImportMockStateFromFile(filePath string) error {
	mockUrl, name, importError := service.CreateAndImportMockState(filePath)
	if importError != nil {
		return importError
	}

	ActiveMocks[name] = mockUrl
	return nil
}
