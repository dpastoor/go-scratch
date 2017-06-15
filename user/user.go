package main

import (
	"fmt"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(fmt.Errorf("error with user %s", err))
	}
	fmt.Println(usr.HomeDir)
}
