package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"io"
	"io/ioutil"
	"os"
	utils "github.com/dpastoor/goutils"
	"strings"
)

func main() {
	fmt.Println("--------- read lines --------")
	fmt.Println(ReadLines("file_with_newline.txt"))
	fmt.Println("with newline")
 checkFileHash("file_with_newline.txt", "newline.txt", false)
	fmt.Println("no newline")
 checkFileHash("file_with_no_newline.txt", "no_newline.txt", false)
	fmt.Println("--------- direct copy --------")
	fmt.Println(ReadLines("file_with_newline.txt"))
	fmt.Println("with newline")
	checkFileHash("file_with_newline.txt", "newline.txt", true)
	fmt.Println("no newline")
	checkFileHash("file_with_no_newline.txt", "no_newline.txt", true)
}

func getHash(p string) string {
	f, err := os.Open(p)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}
//ReadLines reads lines for a file at a given path
func ReadLines(path string) ([]string, error) {
	inFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		fmt.Println(scanner.Bytes())
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func checkFileHash(p string, newPath string, copy bool) {
	fs := afero.NewOsFs()
	fs.Remove(newPath)
	fmt.Println("original hash: " + getHash(p))
	//Get the lines of the file
	if copy {
		input, err := ioutil.ReadFile(p)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = ioutil.WriteFile(newPath, input, 0644)
		if err != nil {
			fmt.Println("Error creating", newPath)
			fmt.Println(err)
			return
		}
	} else {
		sourceLines, err := utils.ReadLines(p)

		if err != nil {
			panic("Unable to read the contents of " + p)
		}

		//We'll use stats for setting the mode of the target file to make sure perms are the same
		stats, err := fs.Stat(p)

		if err != nil {
			panic(err)
		}


		//Write the file contents
		fileContents := strings.Join(sourceLines, "\n")
		afero.WriteFile(fs, newPath, []byte(fileContents), stats.Mode())
	}
	fmt.Println("new hash: " + getHash(newPath))
}
