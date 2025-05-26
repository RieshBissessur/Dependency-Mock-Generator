package main

import (
	"context"
	"mock-generator/managers"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestWithRedis(t *testing.T) {
	t.Log("Starting Redis container test...")
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "redis:latest",
		ExposedPorts: []string{"6379/tcp"},
		WaitingFor:   wait.ForLog("Ready to accept connections"),
	}

	redisC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)

	defer func() {
		if err := redisC.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	}()

	// Get the container's mapped port
	mappedPort, err := redisC.MappedPort(ctx, "6379/tcp")
	require.NoError(t, err)

	// Get the container's host
	host, err := redisC.Host(ctx)
	require.NoError(t, err)

	// Print connection details for verification
	t.Logf("Redis is available on %s:%s", host, mappedPort.Port())
	t.Log("Test completed successfully")
}

func TestSetupMock(t *testing.T) {
	// Act
	managers.RunSetup("test-setup.yaml")

	// Assert
	value, contains := managers.ActiveMocks["TestMock"]

	assert.Equal(t, contains, true)
	assert.Contains(t, value, "http://localhost")
}

func TestImportMock(t *testing.T) {
	// Act
	managers.ImportMockStateFromFile("mappings/Mock-2025-05-26T15:50:31+02:00.json")

	// Assert
	value, contains := managers.ActiveMocks["Mock"]

	assert.Equal(t, contains, true)
	assert.Contains(t, value, "http://localhost")
}
