package provider

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/rieshbissessur/dependency-mock-generator/internal/model"
	"github.com/rieshbissessur/dependency-mock-generator/internal/repository"
	"gopkg.in/yaml.v3"
)

func GetSetupFromFile(filePath string) (model.Setup, error) {
	fileBytes, fileReadError := repository.ReadFileContentAsBytes(filePath)
	if fileReadError != nil {
		return model.Setup{}, fileReadError
	}

	var setup model.Setup
	var raw any

	if marshalError := yaml.Unmarshal(fileBytes, &raw); marshalError != nil {
		fmt.Println("Error unmarshaling setup file bytes:", marshalError)
		return model.Setup{}, marshalError
	}

	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: &setup})
	if decoderError := decoder.Decode(raw); decoderError != nil {
		fmt.Println("Error decoding setup interface: ", decoderError)
		return model.Setup{}, decoderError
	}

	return setup, nil
}
