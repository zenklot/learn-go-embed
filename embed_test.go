package learngoembed

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed version.txt
var version string

func Test_embed(t *testing.T) {
	fmt.Println(version)
}
