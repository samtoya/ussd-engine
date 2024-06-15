package parsers

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"ussd-gateway/pkg/data/screens"
)

type ScreenParser interface {
	Parse(string) (map[string]*screens.Screen, error)
	ParseDefault() (map[string]*screens.Screen, error)
}

type yamlScreenParser struct {
}

func (parser *yamlScreenParser) ParseDefault() (map[string]*screens.Screen, error) {
	obj := make(map[string]*screens.Screen)

	file, err := os.ReadFile("./screen.yaml")
	if err != nil {
		log.Fatalln("failed to read screen.yaml", err)
	}

	if err := yaml.Unmarshal(file, obj); err != nil {
		log.Printf("Unmarshal: %v\n", err)
	}

	isValid, err := screens.ValidateScreenEntry(obj)
	if err != nil {
		log.Println("failed to validate screen", err)
		os.Exit(4)
	}

	if !isValid {
		log.Println("screen types failed validation")
		os.Exit(4)
	}

	log.Println("screen.yaml has been parsed = ", obj)

	return obj, nil
}

func (parser *yamlScreenParser) Parse(s string) (map[string]*screens.Screen, error) {
	obj := make(map[string]*screens.Screen)

	file, err := os.ReadFile(s)
	if err != nil {
		log.Fatalln("failed to read screen.yaml", err)
	}

	if err := yaml.Unmarshal(file, obj); err != nil {
		log.Printf("Unmarshal: %v\n", err)
	}

	isValid, err := screens.ValidateScreenEntry(obj)
	if err != nil {
		log.Println("failed to validate screen", err)
		os.Exit(4)
	}

	if !isValid {
		log.Println("screen types failed validation")
		os.Exit(4)
	}

	log.Println("screen.yaml has been parsed = ", obj)

	return obj, nil
}

func NewYamlScreenParser() ScreenParser {
	return &yamlScreenParser{}
}
