package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	var datas []int = []int{1, 2, 3}

	l1 := &ListNode{datas[0], nil}
	l1Node := l1

	for i := 1; i < len(datas); i++ {
		l1Node.Next = &ListNode{datas[i], nil}
		//从新赋值
		l1Node = l1Node.Next

	}

	var data2 []int = []int{5, 7}

	l2 := &ListNode{data2[0], nil}
	l2Node := l2

	for i := 1; i < len(data2); i++ {
		l2Node.Next = &ListNode{data2[i], nil}
		//从新赋值
		l2Node = l2Node.Next

	}
	fmt.Println(l1)
	l := mergeTwoLists(l1, l2)
	fmt.Println(l.Val, l.Next.Val, l.Next.Next.Val, l.Next.Next.Next.Val,
		l.Next.Next.Next.Next.Val,
		l.Next.Next.Next.Next.Next.Val,
	)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	info := make(map[int]int)
	var maxvalue int
	//基数排序方式将列表转换成数组
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	for {
		if l1 != nil {
			info[l1.Val] += 1
			if l1.Val > maxvalue {
				maxvalue = l1.Val
			}
			l1 = l1.Next
		} else {
			break
		}
	}
	for {
		if l2 != nil {
			info[l2.Val] += 1
			if l2.Val > maxvalue {
				maxvalue = l2.Val
			}
			l2 = l2.Next
		} else {
			break
		}
	}

	lastArray := make([]int, maxvalue+1)

	for k, v := range info {
		fmt.Println(k, v)
		lastArray[k] = v
	}

	//循环数组给出列表
	l := &ListNode{0, nil}
	lNode := l
	for k, v := range lastArray {
		if v != 0 {
			for i := 0; i < v; i++ {
				if lNode.Val == 0 {
					l.Val = k
				} else {
					lNode.Next = &ListNode{k, nil}
					lNode = lNode.Next
				}
			}
		}
	}

	return l
}
