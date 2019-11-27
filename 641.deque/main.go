package main

import (
	"fmt"
)

/*
设计实现双端队列。
你的实现需要支持以下操作：

MyCircularDeque(k)：构造函数,双端队列的大小为k。
insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
isEmpty()：检查双端队列是否为空。
isFull()：检查双端队列是否满了。

所有值的范围为 [1, 1000]
操作次数的范围为 [1, 1000]
请不要使用内置的双端队列库。
*/

func main() {
	fmt.Println("ok")

	obj := Constructor(3)
	param_1 := obj.InsertFront(1)
	param_2 := obj.InsertLast(2)
	//param_3 := obj.DeleteFront()
	//param_4 := obj.DeleteLast()
	param_5 := obj.GetFront()
	param_6 := obj.GetRear()
	param_7 := obj.IsEmpty()
	param_8 := obj.IsFull()

	fmt.Println(param_1)
	fmt.Println(param_2)
	//fmt.Println(param_3)
	//fmt.Println(param_4)
	fmt.Println(param_5)
	fmt.Println(param_6)
	fmt.Println(param_7)
	fmt.Println(param_8)

}

type MyCircularDeque struct {
	info []int
	head int
	tail int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {

	return MyCircularDeque{
		info: make([]int, k+1, k+1),
		head: 0,
		tail: 0,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	// 队头进元素，队头指针减 1（想象一个圆环，往左侧移动一位）
	// 因是原型队列，index不应该超出总长度或出现负值，这里采用%的方式控制
	this.head = (this.head - 1 + len(this.info)) % (len(this.info))
	this.info[this.head] = value
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.info[this.tail] = value
	//对位进元素，队尾指针加 1
	this.tail = (this.tail + 1) % len(this.info)

	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.info[this.head] = 0
	this.head = (this.head + 1) % len(this.info)
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.info[this.tail] = 0
	this.tail = (this.tail - 1 + len(this.info)) % len(this.info)
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.info[this.head]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	//获取尾部队列是需要往左移动一位
	return this.info[(this.tail-1+len(this.info))%len(this.info)]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if this.head == this.tail {
		return true
	}
	return false
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {

	if this.head == (this.tail+1)%len(this.info) {
		return true
	}
	return false
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
