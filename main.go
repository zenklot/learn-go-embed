package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed logo.png
var logo []byte

//go:embed files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := ioutil.WriteFile("new_logo.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dir, _ := path.ReadDir("files")
	for _, entry := range dir {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println("Content:", string(content))
		}
	}
}
