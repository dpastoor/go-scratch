package main

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"regexp"
)

type OsRelease struct {
	Name            string `toml:"NAME"`
	Version         string `toml:"VERSION"`
	Id              string `toml:"ID"`
	IdLike          string `toml:"ID_LIKE"`
	LtsRelease      string
	PrettyName      string `toml:"PRETTY_NAME"`
	VersionId       string `toml:"VERSION_ID"`
	VersionCodename string `toml:"VERSION_CODENAME"`
	UbuntuCodename  string `toml:"UBUNTU_CODENAME"`
}

func main() {
	re := regexp.MustCompile(`(.*?=)([^"].*)`)
	osRel, err := ioutil.ReadFile("os-release")
	if err != nil {
		panic(err)
	}
	fixedConfig := re.ReplaceAll(osRel, []byte("${1}\"${2}\""))
	osr := OsRelease{}
	err = toml.Unmarshal(fixedConfig, &osr)
	if err != nil {
		panic(err)
	}
	fmt.Println(osr)
}
