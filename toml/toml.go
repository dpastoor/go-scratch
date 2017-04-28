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

func main() {
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
