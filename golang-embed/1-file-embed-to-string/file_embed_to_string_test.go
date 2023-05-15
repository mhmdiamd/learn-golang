package fileembed

import (
	_ "embed"
	"fmt"
	"testing"
)

// Cara embed file adalah dengan menggunaakan kode di dalam komentar

//go:embed version.txt
var version string

func TestFileEmbed(t *testing.T) {
	fmt.Println(version)
}
