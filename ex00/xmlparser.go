package main

import (
	"encoding/json"
	"encoding/xml"
	"os"
)

type XMLParserWriter struct{}

func (x XMLParserWriter) Parse(path string) (*Recipes, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var recipes Recipes
	if err := xml.Unmarshal(fileBytes, &recipes); err != nil {
		return nil, err
	}
	return &recipes, nil
}

func (x XMLParserWriter) Write(recipes *Recipes) error {
    jsonData, err := json.MarshalIndent(recipes, "", "    ")
    if err != nil {
        return err
    }
    err = os.WriteFile("recipeFromJSON.xml", jsonData, 0644)
    if err != nil {
        return err
    }
    return nil
}
