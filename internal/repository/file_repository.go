package repository

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadFileContentAsString(fullFilePath string) (string, error) {
	content, fileReadError := os.ReadFile(fullFilePath)
	if fileReadError != nil {
		return "", fmt.Errorf("failed to read file %s: %w", fullFilePath, fileReadError)
	}

	return string(content), nil
}

func ReadFileContentAsBytes(fullFilePath string) ([]byte, error) {
	content, fileReadError := os.ReadFile(fullFilePath)
	if fileReadError != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", fullFilePath, fileReadError)
	}

	return content, nil
}

func WriteFileContent(directory string, fileName string, content string) error {
	directoryError := CreateDirectory(directory)
	if directoryError != nil {
		return directoryError
	}

	fileWriteError := os.WriteFile(fmt.Sprintf("%s/%s", directory, fileName), []byte(content), 0644)
	if fileWriteError != nil {
		return fmt.Errorf("Failed to write file %s/%s: %w", directory, fileName, fileWriteError)
	}

	return nil
}

func WriteFileContentFromAny(directory string, fileName string, content any) error {
	directoryError := CreateDirectory(directory)
	if directoryError != nil {
		return directoryError
	}

	jsonData, err := json.Marshal(content)
	if err != nil {
		return err
	}

	fileWriteError := os.WriteFile(fmt.Sprintf("%s/%s", directory, fileName), []byte(jsonData), 0644)
	if fileWriteError != nil {
		return fmt.Errorf("Failed to write file %s/%s: %w", directory, fileName, fileWriteError)
	}

	return nil
}

func CreateDirectory(directoryPath string) error {
	if _, err := os.Stat(directoryPath); os.IsExist(err) || err == nil {
		return nil
	}

	makeDirErr := os.Mkdir(directoryPath, 0755)
	if makeDirErr != nil {
		return makeDirErr
	}

	return nil
}
