package main

import (
	"fmt"
	"os"
)

/*
给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。

说明:
初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
示例:
输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3
输出: [1,2,2,3,5,6]
*/
func main() {
	var m int = 3
	var n int = 5
	var d1 []int = []int{1, 3, 4, 0, 0, 0, 0, 0}
	var d2 []int = []int{1, 3, 6, 8, 9}
	//merge(d1, m, d2, n)

	nums1 := append(d1[0:m], d2[0:n]...)
	usemaopao(nums1)
}

/*
* 使用倒序方式 m+n-1
 */
func merge(nums1 []int, m int, nums2 []int, n int) {

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
	var isstop bool = true
	for i := 0; i < len(nums1)-1; i++ {
		fmt.Println("start=", nums1)
		isstop = true
		for j := 0; j < len(nums1)-i-1; j++ {
			if nums1[j] > nums1[j+1] {
				nums1[j], nums1[j+1] = nums1[j+1], nums1[j]
				isstop = false
			}
		}
		if isstop == true {
			fmt.Println("没有冒泡输出，进行终止")
			break
		}
		fmt.Println("end=", nums1)
	}
	fmt.Println(nums1)
	os.Exit(0)
	fmt.Println(nums1)
}
