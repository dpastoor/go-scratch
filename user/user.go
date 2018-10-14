package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(fmt.Errorf("error with user %s", err))
	}
	fmt.Println("HOMEDIR:", usr.HomeDir)
	fmt.Println("Name:", usr.Name)
	fmt.Println("Username:", usr.Username)
	cdir, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}
	fmt.Println("user cache dir: ", cdir)
}
