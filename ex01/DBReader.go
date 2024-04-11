package main

type DBReader interface {
    Parse(path string) (*Recipes, error)
    Write(recipes *Recipes) error
}
