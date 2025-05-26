package managers

import (
	"mock-generator/providers"
	"mock-generator/services"
)

var activeContainers = make(map[string]string)
var activeMocks = make(map[string]string)

func RunSetup(setupPathFile string) error {

	setup, setupErr := providers.GetSetupFromFile(setupPathFile)
	if setupErr != nil {
		return nil
	}

	for _, container := range setup.Containers {
		containerUrl, err := services.SetupContainer(container)
		if err != nil {
			continue
		}

		activeContainers[container.Name] = containerUrl
	}

	for _, wireMock := range setup.Mocks {
		mockUrl, err := services.SetupMock(wireMock)
		if err != nil {
			continue
		}

		activeMocks[wireMock.Name] = mockUrl
	}

	return nil
}

func ExportMockStates() error {
	for name, wireMock := range activeMocks {
		exportError := services.ExportMockState(name, wireMock)
		if exportError != nil {
			return exportError
		}
	}

	return nil
}

func ImportMockStateFromFile(filePath string) error {
	mockUrl, name, importError := services.ImportMockState(filePath)
	if importError != nil {
		return importError
	}

	activeMocks[name] = mockUrl
	return nil
}
