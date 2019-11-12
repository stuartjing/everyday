package main

import (
	"fmt"
)

func main() {

	var a string = "aacc"
	var b string = "ccac"
	bool1 := isAnagram2(a, b)
	fmt.Println(bool1)
}

//将所有内容放到一个大数组中，第一个字符串有这个值，就+，第二个字符串有就减，
//最后结果都为0怎返回true
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s1 := []rune(s)
	t1 := []rune(t)

	str := make(map[rune]int, len(s))

	for i := range s1 {
		str[s1[i]]++
		str[t1[i]]--
	}

	for _, v := range str {
		if v != 0 {
			return false
		}
	}
	return true

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
