package providers

import (
	"fmt"
	"mock-generator/repositories"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/yaml.v3"
)

type Setup struct {
	Containers []Container
	Mocks      []Mock
}

type Container struct {
	Name  string
	Image string
	Tag   string
	Port  string
}

type Mock struct {
	Name string
	Type int
	File string
}

func GetSetupFromFile(filePath string) (Setup, error) {
	fileBytes, fileReadError := repositories.ReadFileContentAsBytes(filePath)
	if fileReadError != nil {
		return Setup{}, fileReadError
	}

	var setup Setup
	var raw any

	if marshalError := yaml.Unmarshal(fileBytes, &raw); marshalError != nil {
		fmt.Println("Error unmarshaling setup file bytes:", marshalError)
		return Setup{}, marshalError
	}

	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: &setup})
	if decoderError := decoder.Decode(raw); decoderError != nil {
		fmt.Println("Error decoding setup interface: ", decoderError)
		return Setup{}, decoderError
	}

	return setup, nil
}
