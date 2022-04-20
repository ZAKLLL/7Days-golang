package gee

import (
	"testing"
)

func Test(t *testing.T) {
	root := node{}
	for _, item := range []string{"/", "/hello", "/hi/:user/do", "/hi/:xxx/do1", "/hi/t1", "/hi/*", "/hi/**"} {
		root.insert(item, parsePattern(item), 0)
	}
	println(root.search([]string{"hi", "xxx", "do1x"}, 0).pattern)

	println(root.search([]string{""}, 0).pattern)

	println(root.search([]string{"hi", "xx", "do"}, 0).pattern)

	println(root.search([]string{"hi", "t1"}, 0).pattern)

	println(root.search([]string{"hi", "t2"}, 0).pattern)

	println(root.search([]string{"hi", "t3", "t4"}, 0).pattern)
}
