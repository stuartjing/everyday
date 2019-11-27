package main

import (
	"fmt"
)

/*
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
示例：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

*/

func main() {

	var datas []int = []int{1, 2, 6}
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
	//	fmt.Println(l1)
	l := mergeTwoLists2(l1, l2)

	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	newList := head
	for l1 != nil && l2 != nil {
		//对比遍历到当前的节点
		if l1.Val < l2.Val {
			//将l1数据加入赋给newlist的next
			newList.Next = l1
			//将指针往下移动
			l1 = l1.Next
		} else {
			newList.Next = l2
			l2 = l2.Next
		}
		newList = newList.Next
	}

	if l1 != nil {
		newList.Next = l1
	} else {
		newList.Next = l2
	}
	return head.Next
}

/*
递归方式
*/
func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val > l2.Val {
		l2.Next = mergeTwoLists2(l1, l2.Next)
		return l2
	} else {
		l1.Next = mergeTwoLists2(l1.Next, l2)
		return l1
	}
}
