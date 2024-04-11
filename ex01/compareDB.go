package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strings"
)

type CompareDB interface {
	Compare(*Recipes, *Recipes)
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

func printError(err string) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	if len(os.Args) == 1 {
		printError("./compareDB --old original_database.xml --new stolen_database.json")
	} else {
		args := os.Args[1:]
		if args[0] != "--old" {
			printError("./compareDB --old original_database.xml --new stolen_database.json")
		} else if args[2] != "--new" {
			printError("./compareDB --old original_database.xml --new stolen_database.json")
		}
		if len(args) < 4 {
			printError("./compareDB --old original_database.xml --new stolen_database.json")
		}
		var compareDB CompareDB

		var parseOldDB DBReader
		var parseNewDB DBReader

		if !strings.HasSuffix(args[1], ".json") && !strings.HasSuffix(args[1], ".xml") {
			printError("Error: no such file in arg 2")
		} else {
			if strings.HasSuffix(args[1], ".json") {
				parseOldDB = JSONParserWriter{}
			} else {
				parseOldDB = XMLParserWriter{}
			}
		}
		if !strings.HasSuffix(args[3], ".json") && !strings.HasSuffix(args[3], ".xml") {
			printError("Error: no such file in arg 4")
		} else {
			if strings.HasSuffix(args[3], ".json") {
				parseNewDB = JSONParserWriter{}
			} else {
				parseNewDB = XMLParserWriter{}
			}
		}
		parsedOld, err := parseOldDB.Parse(args[1])
		if err != nil {
			printError("Error: cant parse file in args 1")
		}
		parsedNew, err := parseNewDB.Parse(args[3])
		if err != nil {
			printError("Error: cant parse file in args 1")
		}
		compareDB = CompareDBs{}
		compareDB.Compare(parsedOld, parsedNew)
	}
}