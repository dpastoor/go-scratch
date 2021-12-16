package main

import (
	"fmt"

	"github.com/metrumresearchgroup/command"
)

func main() {
	cmd := command.New("pwd")
	_ = cmd.Impersonate("devinp", false)
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
