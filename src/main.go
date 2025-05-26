package main

import (
	"fmt"
	"mock-generator/managers"
)

func main() {

	err := managers.RunSetup("setup.yaml")

	if err != nil {
		fmt.Println("Error with setup setup:", err)
		return
	}

	//managers.ExportMockStates()

	managers.ImportMockStateFromFile("/Users/rieshbissessur/Downloads/Git/Mock-Generator/src/mappings/Mock-2025-05-26T15:50:31+02:00.json")

	for true {
	}
}
