package main

import (
	"fmt"
)

func main() {
	/*
		//二分查找
		var arr []int = []int{1, 2, 3, 4, 5, 6, 7, 8}
		var key int = 7
		isOk := BinarySearch(arr, key)
		fmt.Println(isOk)
		isOk2 := BinarySearchRec(arr, 0, len(arr), key)
		fmt.Println(isOk2)
	*/

	//快速排序
	var arr1 []int = []int{6, 3, 3, 4, 11, 6, 13, 8}
	quickSort2(arr1, 0, len(arr1)-1)
	fmt.Println(arr1)

	fibnum := 4
	//斐波那契数列
	fibres := fibonacci(fibnum)
	fmt.Println(fibres)

	//斐波那契数列
	fibres1 := fibonaccibase(fibnum)
	fmt.Println(fibres1)

}
