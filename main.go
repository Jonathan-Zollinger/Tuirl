package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {

	tilNotes := TILDoc{}
	yamlSource, err := os.ReadFile("TIL.yml")

	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = yaml.Unmarshal(yamlSource, &tilNotes)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// ImportYamlNotes("TIL.yml", &tilNotes)
	// fmt.Println("\nprinting tilNotes:\n", tilNotes)
}

type TILDoc struct {
	Title    string `yaml:"title"`
	SubTitle string `yaml:"sub_title"`
}

type Topic struct {
	TopicName string `yaml:"topic-name"`
	Note      string `yaml:"note"`
	Example   string `yaml:"example"`
	// Topics    []Topic `yaml:"sub-topics"`
}
