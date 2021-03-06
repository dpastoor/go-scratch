package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type simpleConfig struct {
	Age        int
	Cats       []string
	Pi         float64 `toml:"pi,omitzero"`
	Perfection []int
}

type nestedConfig struct {
	Top    int
	Nested simpleConfig
}

type model struct {
	Type     string `toml:"type"`
	Language string `toml:"language"`
}

type nestedArray struct {
	Top    int
	Models []model `toml:"models"`
}

func main() {
	nestedArrayTest()
}

func nestedArrayTest() {
	var result nestedArray
	if _, err := toml.DecodeFile("nested_arrays.toml", &result); err != nil {
		fmt.Printf("could not decode")
	}
	fmt.Println(result)
}

func nestedTest() {
	var nested nestedConfig
	if _, err := toml.DecodeFile("nested.toml", &nested); err != nil {
		fmt.Printf("could not decode")
	}

	fmt.Println(nested)

	outputSimple := simpleConfig{
		Age:        29,
		Cats:       []string{"Fred", "George"},
		Pi:         3.1415,
		Perfection: []int{1, 2, 3},
	}

	outputNested := nestedConfig{
		Top:    3,
		Nested: outputSimple,
	}

	f, _ := os.Create("output_nested.toml")
	defer f.Close()

	output := toml.NewEncoder(f)
	err := output.Encode(outputNested)
	if err != nil {
		fmt.Println(fmt.Sprintf("error on encoding: %s", err))
	}
	fmt.Println(outputNested)

}

func simpleTest() {
	var simple simpleConfig
	if _, err := toml.DecodeFile("simple.toml", &simple); err != nil {
		fmt.Printf("could not decode")
	}

	fmt.Println(simple)

	outputSimple := simpleConfig{
		Age:        29,
		Cats:       []string{"Fred", "George"},
		Pi:         3.1415,
		Perfection: []int{1, 2, 3},
	}
	outputSimpleMissing := simpleConfig{
		Age:        29,
		Cats:       []string{"Fred", "George"},
		Perfection: []int{1, 2, 3},
	}
	f, _ := os.Create("output_simple.toml")
	defer f.Close()

	output := toml.NewEncoder(f)
	err := output.Encode(outputSimple)
	if err != nil {
		fmt.Println(fmt.Sprintf("error on encoding: %s", err))
	}
	fmt.Println(outputSimple)

	f2, _ := os.Create("output_simple_missing.toml")
	defer f2.Close()

	output2 := toml.NewEncoder(f2)
	err = output2.Encode(outputSimpleMissing)
	if err != nil {
		fmt.Println(fmt.Sprintf("error on encoding: %s", err))
	}

}
