package main

import "fmt"

type Name struct {
	name string
}

type generateName func() string

func main() {
	fmt.Printf("Product result = %s\n", genTrickyPayload(trickyGenerator))
}

func genTrickyPayload(fn generateName) Name {
	genName := fn()
	result := Name{
		name: genName,
	}
	return result
}

func trickyGenerator() string {
	return "generated-name-"
}
