package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

type DBReader interface {
    Parse(path string) (*Recipes, error)
    Write(recipes *Recipes) error
}

type Recipes struct {
	XMLName xml.Name `json:"-" xml:"recipes"`
	Cakes []Cake `json:"cake" xml:"cake"`
}

type Cake struct {
	XMLName xml.Name `json:"-" xml:"cake"`
	Name string `json:"name" xml:"name"`
	Time string `json:"time" xml:"stovetime"`
	Ingredients []Ingredient `json:"ingredients" xml:"ingredients>item"`
}

type Ingredient struct {
	Name  string `json:"ingredient_name" xml:"itemname"`
	Count string `json:"ingredient_count" xml:"itemcount"`
	Unit  string `json:"ingredient_unit,omitempty" xml:"itemunit,omitempty"`
}



func main() {
	if len(os.Args) == 1 {
        fmt.Println("Usage: program_name -f filename")
        os.Exit(1)
    } else {
        args := os.Args[1:]
        if args[0] != "-f" {
            fmt.Println("Error: -f flag is missing")
            os.Exit(1)
        }
        if len(args) < 2 {
            fmt.Println("Error: filename is missing")
            os.Exit(1)
        }

        var parserWriter DBReader
        
        if strings.HasSuffix(args[1], ".json") {
            parserWriter = JSONParserWriter{}
        } else if strings.HasSuffix(args[1], ".xml") {
            parserWriter = XMLParserWriter{}
        } else {
            fmt.Println("Error: it is not a JSON or XML file")
            os.Exit(1)
        }

        parsed, err := parserWriter.Parse(args[1])
        if err != nil {
            fmt.Println("Error:", err)
            os.Exit(1)
        }

        if err := parserWriter.Write(parsed); err != nil {
            fmt.Println("Error:", err)
            os.Exit(1)
        }
    }
}