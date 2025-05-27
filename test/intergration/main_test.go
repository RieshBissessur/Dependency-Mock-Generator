package intergration

import (
	"testing"

	"github.com/rieshbissessur/dependency-mock-generator/internal/manager"
	"github.com/stretchr/testify/assert"
)

func TestSetupMock(t *testing.T) {
	// Act
	runError := manager.RunSetup("../../configs/test-mock-setup.yaml")

	// Assert
	assert.Equal(t, runError, nil)
	value, contains := manager.ActiveMocks["TestMock"]

	assert.Equal(t, contains, true)
	assert.Contains(t, value, "http://localhost")
}

func TestSetupContainer(t *testing.T) {
	// Act
	runError := manager.RunSetup("../../configs/test-container-setup.yaml")

	// Assert
	assert.Equal(t, runError, nil)
	_, contains := manager.ActiveContainers["Redis"]

	assert.Equal(t, contains, true)
}

func TestImportMock(t *testing.T) {
	// Act
	importError := manager.ImportMockStateFromFile("../../mappings/Mock-2025-05-26T15:50:31+02:00.json")

	// Assert
	assert.Equal(t, importError, nil)
	value, contains := manager.ActiveMocks["Mock"]

	assert.Equal(t, contains, true)
	assert.Contains(t, value, "http://localhost")
}
