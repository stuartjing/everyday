package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	fmt.Println("快速排序：quickSort")
	temp, i := arr[0], 1
	min, max := 0, len(arr)-1
	for min < max {

		if arr[i] > temp {
			arr[max], arr[i] = arr[i], arr[max]
			max--
		} else {
			arr[min], arr[i] = arr[i], arr[min]
			min++
			i++
		}
	}
	arr[min] = temp
	return quickSort(arr[:min])
	return quickSort(arr[min+1:])
}

func quickSort2(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		//先跟最右边比
		for j >= p && values[j] >= temp {
			j--
		}
		//将当前值跟最右边对比出来的进行调整
		if j >= p {
			values[p] = values[j]
			p = j
		}

		//在跟最左边比
		for i <= p && values[i] <= temp {
			i++
		}
		//将当前值和最左边进行调换
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}

	values[p] = temp
	if p-left > 1 {
		quickSort2(values, left, p-1)
	}
	if right-p > 1 {
		quickSort2(values, p+1, right)
	}
}
