package learngoembed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func Test_embed(t *testing.T) {
	fmt.Println(version)
}

//go:embed logo.png
var logo []byte

func Test_byte(t *testing.T) {
	err := ioutil.WriteFile("new_logo.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
