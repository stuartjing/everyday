package main

import (
	"fmt"
)

func main() {

	var a string = "aacc"
	var b string = "ccac"
	bool1 := isAnagram(a, b)
	fmt.Println(bool1)
}

//排序转换成rune计数方式map，在进行对比是否存在
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	a := make(map[rune]int, len(t))
	for _, v := range s {
		a[v-'a'] += 1
	}

	for _, v1 := range t {
		if a[v1-'a'] == 0 {
			return false
		}
		a[v1-'a'] -= 1
	}

	return true
}
