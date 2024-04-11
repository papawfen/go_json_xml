package main

import (
	"fmt"
)

type CompareDBs struct{}

func (c CompareDBs) Compare(old *Recipes, new *Recipes) {
	set := make(map[string]bool)
	setNew := make(map[string]bool)
	for _, cake := range old.Cakes {
		set[cake.Name] = true
	}
	for _, cake := range new.Cakes {
		setNew[cake.Name] = true
	}

	for _, cake := range new.Cakes {
		if !set[cake.Name] {
			fmt.Println("REMOVED cake \"" + cake.Name + "\"")
		}
	}
	for _, cake := range old.Cakes {
		if !setNew[cake.Name] {
			fmt.Println("ADDED cake \"" + cake.Name + "\"")
		}
	}

	for _, cake := range new.Cakes {
		for _, orig := range old.Cakes {
			if cake.Name == orig.Name {
				if cake.Time != orig.Time {
					fmt.Println("CHANGED cooking time for cake \"" + cake.Name + "\" - \"" +
						cake.Time + "\" instead of \"" + orig.Time + "\"")
				}
			}
		}
	}

	for _, cake := range new.Cakes {
		for _, orig := range old.Cakes {
			if cake.Name == orig.Name {
				compareIngredients(orig.Ingredients, cake.Ingredients, cake.Name)
			}
		}
	}

}

func compareIngredients(old []Ingredient, new []Ingredient, cakeName string) {
	oldSet := make(map[string]bool)
	newSet := make(map[string]bool)
	for _, ing := range old {
		oldSet[ing.Name] = true
	}
	for _, ing := range new {
		newSet[ing.Name] = true
	}

	for _, ingredient := range old {
		if !newSet[ingredient.Name] {
			fmt.Println("ADDED ingredient \"" + ingredient.Name + "\"" + "for cake \"" + cakeName + "\"")
		}
	}
	for _, ingredient := range new {
		if !oldSet[ingredient.Name] {
			fmt.Println("REMOVED ingredient \"" + ingredient.Name + "\"" + "for cake \"" + cakeName + "\"")
		}
	}

	for _, original := range old {
		for _, stolen := range new {
			if original.Name == stolen.Name {
				if original.Unit != stolen.Unit {
					if original.Unit == "" && stolen.Unit != "" {
						fmt.Println("REMOVED unit " + stolen.Unit + " for ingredient " + original.Name + " for cake " + cakeName)
					} else if stolen.Unit == "" && original.Unit != "" {
						fmt.Println("ADDED unit " + stolen.Unit + " for ingredient " + original.Name + " for cake " + cakeName)
					} else {
						fmt.Println("CHANGED unit for ingredient " + original.Name + " for cake " + cakeName + " - " + 
									stolen.Unit + " instead of " + original.Unit)
					}
				} else {
					if original.Count != stolen.Count {
						fmt.Println("CHANGED unit count for ingredient " + original.Name + " for cake " + cakeName + " - " + 
									stolen.Count + " instead of " + original.Count)
					}
				}
			}
		}
	}
}
