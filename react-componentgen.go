package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {

	type Configuration struct {
		TemplateFile             string
		TemplateReplacementToken string
		OutDirectory             string
	}

	// Load config
	config, err := os.Open("react-componentgen.conf.json")
	if err != nil {
		fmt.Println("Error opening react-componentgen.conf.json - check file exists.\nError:", err)
		return
	}
	defer config.Close()
	decoder := json.NewDecoder(config)
	configuration := Configuration{}
	decoderErr := decoder.Decode(&configuration)
	if decoderErr != nil {
		fmt.Println("Error in parsing react-componentgen.conf.json - check file.\nError:", decoderErr)
		return
	}

	// Load template file
	componentTemplateFile, err := os.ReadFile(configuration.TemplateFile)

	if err != nil {
		fmt.Println("Error reading component template file. Please check file.\nError:", err)
		return
	}

	var componentTemplate string = string(componentTemplateFile)

	// Get name of desired component
	fmt.Println("Enter name of desired component:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Component name read error. Please try again.")
		return
	}

	var componentName string = strings.TrimSuffix(input, "\r\n")

	// Go through template and replace generic token with desired name
	s := strings.ReplaceAll(componentTemplate, configuration.TemplateReplacementToken, componentName)

	// output to file in directory structure
	outputFile, err := os.Create(configuration.OutDirectory + componentName + ".tsx")
	if err == nil {
		fmt.Println("Writing file...")
		outputFile.Write([]byte(s))
		outputFile.Close()
	} else {
		fmt.Print(err)
		return
	}

	fmt.Printf("%s written.", configuration.OutDirectory+componentName+".tsx")

}
