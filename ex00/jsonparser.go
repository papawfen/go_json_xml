package main

import (
	"encoding/json"
	"encoding/xml"
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
    file, err := os.Create("recipeFromJSON.xml")
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.WriteString(xml.Header); err != nil {
		return err
	}
	
	enc := xml.NewEncoder(file)
	enc.Indent("", "    ")

	if err := enc.EncodeToken(xml.StartElement{Name: xml.Name{Local: "recipes"}}); err != nil {
		return err
	}

	for _, cake := range recipes.Cakes {
		if err := enc.Encode(cake); err != nil {
			return err
		}
	}

	if err := enc.EncodeToken(xml.EndElement{Name: xml.Name{Local: "recipes"}}); err != nil {
		return err
	}

	if err := enc.Flush(); err != nil {
		return err
	}

	return nil
}
