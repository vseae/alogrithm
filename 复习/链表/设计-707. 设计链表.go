package main

type Node struct {
	Val  int
	Next *Node
	Pre  *Node
}
type MyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (linkedList *MyLinkedList) Get(index int) int {
	if linkedList.size == 0 || index < 0 || index >= linkedList.size {
		return -1
	}
	if index == 0 {
		return linkedList.head.Val
	}
	if index == linkedList.size {
		return linkedList.tail.Val
	}
	count := 0
	cur := linkedList.head
	for cur != nil {
		if count == index {
			break
		}
		count++
		cur = cur.Next
	}
	return cur.Val
}

func (linkedList *MyLinkedList) AddAtHead(val int) {
	node := &Node{Val: val}
	if linkedList.head != nil {
		node.Next = linkedList.head
		linkedList.head.Pre = node
		linkedList.head = node

	} else {
		linkedList.head = node
		linkedList.tail = linkedList.head
	}
	linkedList.size++
}

func (linkedList *MyLinkedList) AddAtTail(val int) {
	node := &Node{Val: val}
	if linkedList.tail != nil {
		node.Pre = linkedList.tail
		linkedList.tail.Next = node
		linkedList.tail = node
	} else {
		linkedList.tail = node
		linkedList.head = linkedList.tail
	}
	linkedList.size++
}

func (linkedList *MyLinkedList) AddAtIndex(index, val int) {
	if index > linkedList.size {
		return
	}
	if index <= 0 {
		linkedList.AddAtHead(val)
		return
	}
	if index == linkedList.size {
		linkedList.AddAtTail(val)
		return
	}
	cur := linkedList.head
	count := 0
	for cur != nil && count < index {
		count++
		cur = cur.Next
	}
	linkedList.size++
	node := &Node{Val: val}
	node.Next = cur
	node.Pre = cur.Pre
	cur.Pre.Next = node
	cur.Pre = node
}

func (linkedList *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= linkedList.size || linkedList.size == 0 {
		return
	}
	if index == 0 {
		linkedList.head = linkedList.head.Next
	} else if index == linkedList.size-1 {
		linkedList.tail = linkedList.tail.Pre
	} else {
		cur := linkedList.head
		count := 0
		for cur != nil && count < index {
			count++
			cur = cur.Next
		}
		cur.Pre.Next = cur.Next
		cur.Next.Pre = cur.Pre
	}
	linkedList.size--
}
