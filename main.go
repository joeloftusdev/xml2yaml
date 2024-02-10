package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type CD struct { // This will need to be changed to reflect your files.
	Title   string `xml:"TITLE"`
	Artist  string `xml:"ARTIST"`
	Country string `xml:"COUNTRY"`
	Company string `xml:"COMPANY"`
	Price   string `xml:"PRICE"`
	Year    string `xml:"YEAR"`
}

type Catalog struct {
	XMLName xml.Name `xml:"CATALOG"`
	CDs     []CD     `xml:"CD"`
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal("go run . 'input.xml' 'output.yaml'")
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	xmlFile, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error opening XML file: %v", err)
	}
	defer xmlFile.Close()

	byteValue, err := io.ReadAll(xmlFile)
	if err != nil {
		log.Fatalf("Error reading XML file: %v", err)
	}

	if len(byteValue) == 0 {
		log.Fatal("Empty XML file")
	}

	var catalog Catalog
	err = xml.Unmarshal(byteValue, &catalog)
	if err != nil {
		log.Fatalf("Error parsing XML: %v", err)
	}

	yamlData, err := yaml.Marshal(&catalog)
	if err != nil {
		log.Fatalf("Error marshalling YAML: %v", err)
	}

	output, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer output.Close()

	_, err = output.Write(yamlData)
	if err != nil {
		log.Fatalf("Error writing YAML data: %v", err)
	}

	fmt.Printf("XML file %s converted to YAML and saved to %s\n", inputFile, outputFile)
}
