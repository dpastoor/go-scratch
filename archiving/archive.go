package main

import (
	"fmt"
	"os"

	"github.com/mholt/archiver"
)

func main() {
	err := archiver.TarGz.Make("files.tar.gz", []string{"file1.txt", "file2.txt"})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("successfully targz'd")
}
