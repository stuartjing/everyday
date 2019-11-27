package main

import "fmt"

/*
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
示例 1:
输入: [1,2,3,4,5,6,7] 和 k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右旋转 1 步: [7,1,2,3,4,5,6]
向右旋转 2 步: [6,7,1,2,3,4,5]
向右旋转 3 步: [5,6,7,1,2,3,4]
示例 2:
输入: [-1,-100,3,99] 和 k = 2
输出: [3,99,-1,-100]
解释:
向右旋转 1 步: [99,-1,-100,3]
向右旋转 2 步: [3,99,-1,-100]
说明:
尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
要求使用空间复杂度为 O(1) 的 原地 算法。
*/

func main() {

	var arr []int = []int{1, 2, 3, 4, 5, 6, 7}
	rotate1t(arr, 3)
	fmt.Println(arr)
}

func rotate1(nums []int, k int) {
	k = k % len(nums)
	var nums1 []int = make([]int, len(nums))
	if k != 0 && len(nums) != 1 {
		kk := len(nums) - k
		a1 := nums[:kk]
		a2 := nums[kk:]
		nums1 = append(a2, a1...)
		for tt, v := range nums1 {
			nums[tt] = v
		}
	}
}

func rotate(nums []int, k int) {
	length := len(nums)

	var nums1 []int = make([]int, length)
	// var j int = 0
	// var t int = 0

	if length != 1 {

		k = k % length

		for i := 0; i < length; i++ {
			if k > i {
				nums1[i] = nums[length-k+i]
			} else {
				nums1[i] = nums[i-k]
			}
		}
		fmt.Println(nums1)

		for tt, v := range nums1 {
			nums[tt] = v
		}
	}
	fmt.Println(nums)
}
