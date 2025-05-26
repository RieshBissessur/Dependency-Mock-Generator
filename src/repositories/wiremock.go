package repositories

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mock-generator/models"
	"net/http"
	"strconv"

	"github.com/wiremock/go-wiremock"
)

func AddGetMappingModel(url string, path string, statusCode int64, response map[string]any) error {
	wiremockClient := wiremock.NewClient(url)

	basicAuthStub := wiremock.
		Get(wiremock.URLMatching(path)).
		WillReturnResponse(wiremock.NewResponse().
			WithStatus(statusCode).
			WithJSONBody(response))

	stubError := wiremockClient.StubFor(basicAuthStub)
	if stubError != nil {
		return fmt.Errorf("failed to stub mappings: %w", stubError)
	}

	return nil
}

func AddPostMappingModel(url string, path string, statusCode int64, response map[string]any) error {
	wiremockClient := wiremock.NewClient(url)

	basicAuthStub := wiremock.Post(wiremock.URLMatching(path)).
		WillReturnResponse(wiremock.NewResponse().
			WithStatus(statusCode).
			WithJSONBody(response))

	stubError := wiremockClient.StubFor(basicAuthStub)
	if stubError != nil {
		return fmt.Errorf("failed to stub mappings: %w", stubError)
	}

	return nil
}

func ClearAllMappingModels(url string) error {
	wiremockClient := wiremock.NewClient(url)

	clearError := wiremockClient.Reset()
	if clearError != nil {
		return fmt.Errorf("failed to clear mappings: %w", clearError)
	}

	return nil
}

func GetAllMappingModels(url string) (string, error) {
	response, getError := http.Get(fmt.Sprintf("%s/__admin/mappings", url))
	if getError != nil {
		return "", fmt.Errorf("Failed to get mappings: %w", getError)
	}

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Failed to get mappings with status code: %s", strconv.Itoa(response.StatusCode))
	}

	defer response.Body.Close()
	body, readingError := io.ReadAll(response.Body)
	if readingError != nil {
		return "", fmt.Errorf("Failed to read response: %w", readingError)
	}

	return string(body), nil

}

func ImportMappingModels(url string, mapping models.Mapping) error {
	postUrl := fmt.Sprintf("%s/__admin/mappings", url)
	jsonBytes, err := json.Marshal(mapping)
	if err != nil {
		return fmt.Errorf("failed to serialize mapping: %w", err)
	}

	response, postError := http.Post(postUrl, "application/json", bytes.NewReader(jsonBytes))
	if postError != nil {
		return fmt.Errorf("Failed to read response: %w", postError)
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("Failed to get mappings with status code: %s", strconv.Itoa(response.StatusCode))
	}

	return nil
}
