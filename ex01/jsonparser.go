package main

import (
	"encoding/json"
	"os"
)

type JSONParserWriter struct{}

func (j JSONParserWriter) Parse(path string) (*Recipes, error) {
    fileBytes, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }
    var recipes Recipes
    if err := json.Unmarshal(fileBytes, &recipes); err != nil {
        return nil, err
    }
    return &recipes, nil
}

func (j JSONParserWriter) Write(recipes *Recipes) error {
    jsonData, err := json.MarshalIndent(recipes, "", "    ")
    if err != nil {
        return err
    }
    err = os.WriteFile("recipeFromXML.json", jsonData, 0644)
    if err != nil {
        return err
    }
    return nil
}
