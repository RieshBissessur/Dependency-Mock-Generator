package main

import (
	"fmt"

	"github.com/rieshbissessur/dependency-mock-generator/internal/manager"
)

func main() {

	setupError := manager.RunSetup("../configs/setup.yaml")

	if setupError != nil {
		fmt.Println("Error with setup:", setupError)
		return
	}

	// Export Example
	//exportError := manager.ExportMockStates()
	//if exportError != nil {
	//	fmt.Println("Error with export:", exportError)
	//	return
	//}

	// Import Example
	//importError := manager.ImportMockStateFromFile("../../mappings/Mock-2025-05-26T15:50:31+02:00.json")
	//if importError != nil {
	//	fmt.Println("Error with import:", importError)
	//	return
	//}

	for true {
	}
}
