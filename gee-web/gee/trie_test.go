package gee

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	root := node{}
	for _, item := range []string{"/", "/hello", "/hi/:user/do", "/hi/*/v1", "/hi/*", "/hi/*/:lang/zzz"} {
		root.insert(item, parsePattern(item), 0)
	}
	fmt.Print(root)
}
