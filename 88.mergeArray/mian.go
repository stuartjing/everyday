package main

import (
	"fmt"
)

func main() {

	var d1 []int = []int{1, 3, 4, 0, 0, 0, 0, 0}
	var d2 []int = []int{1, 3, 6, 8, 9}
	merge(d1, 3, d2, 5)
}

/*
* 使用倒序方式 m+n-1
 */
func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) == 0 || len(nums2) == 0 {

	}

	for {
		if m == 0 || n == 0 {
			break
		}
		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
	for {
		if n == 0 {
			break
		}
		nums1[m+n-1] = nums2[n-1]
		n--
	}
	fmt.Println(nums1)

	//nums1 = append(nums1[0:m], nums2[0:n]...)
	//usemaopao(nums1)
}

/*
*冒泡方式
 */
func usemaopao(nums1 []int) {

	for i := 0; i < len(nums1)-1; i++ {

		for j := 0; j < len(nums1)-i-1; j++ {
			if nums1[j] > nums1[j+1] {
				nums1[j], nums1[j+1] = nums1[j+1], nums1[j]
			}
		}
	}
	fmt.Println(nums1)
}
