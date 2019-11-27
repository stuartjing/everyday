package main

//使用递归方式
//二分查找
func BinarySearchRec(arr []int, min, max, key int) int {

	if min >= max {
		return -1
	}
	middle := (min + max) / 2

	if key == arr[middle] {
		return middle
	} else if key > arr[middle] {
		return BinarySearchRec(arr, middle+1, max, key)
	} else {
		return BinarySearchRec(arr, min, middle, key)
	}

	return -1
}

//二分查找
func BinarySearch(arr []int, key int) int {
	pos := -1
	min, max := 0, len(arr)

	for {
		if min >= max {
			break
		}
		middle := (max + min) / 2
		//fmt.Println(min, max)
		if key == arr[middle] {
			pos = middle
			break
		} else if key > arr[middle] {
			min = middle + 1
		} else if key < arr[middle] {
			max = middle
		}
	}
	return pos
}
