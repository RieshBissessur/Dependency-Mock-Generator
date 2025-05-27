package main

import (
	"fmt"

	"github.com/rieshbissessur/dependency-mock-generator/internal/manager"
)

func main() {

	err := manager.RunSetup("./config/setup.yaml")

	if err != nil {
		fmt.Println("Error with setup setup:", err)
		return
	}

	//manager.ExportMockStates()

	//manager.ImportMockStateFromFile("mappings/Mock-2025-05-26T15:50:31+02:00.json")

	for true {
	}
}
