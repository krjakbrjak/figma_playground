package main

import (
	"flag"
	"fmt"
)

func main() {
	var fileKey string
	var componentId string
	var apiKey string

	// Parsing flags
	flag.StringVar(&fileKey, "fileKey", "", "Specify Figma file key")
	flag.StringVar(&componentId, "componentId", "", "Specify Figma component ID")
	flag.StringVar(&apiKey, "apiKey", "", "Specify Figma API key")

	// Parse command-line arguments
	flag.Parse()

	figma := Figma{
		FILE_KEY:     fileKey,
		COMPONENT_ID: componentId,
		API_KEY:      apiKey,
	}

	component, err := figma.FetchComponent()
	if err != nil {
		fmt.Println("Error fetching Figma component:", err)
		return
	}

	qml, qmlErr := GenerateQml(component)
	if qmlErr != nil {
		fmt.Println("Error generating QML component:", qmlErr)
		return
	}

	fmt.Println(qml)
}
