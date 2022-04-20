package gee

import (
	"testing"
)

func Test(t *testing.T) {
	root := node{}
	for _, item := range []string{"/", "/hello", "/hi/:user/do", "/hi/:xxx/do1", "/hi/t1", "/hi/*", "/hi/**"} {
		root.insert(item, parsePattern(item), 0)
	}
	search := root.search([]string{"hi", "xxx", "vvv"}, 0)
	println(search.pattern)
}

func Test_parsePattern(t *testing.T) {
	//for _, s := range parsePattern("/hi/:user/do") {
	//	fmt.Println(s)
	//}
	//fmt.Println("-----")
	//
	//for _, s := range parsePattern("/hi/*/v1") {
	//	fmt.Println(s)
	//}
	a := node{}.pattern
	println(a == "")
}
