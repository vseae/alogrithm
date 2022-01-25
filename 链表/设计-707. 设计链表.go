package main

type MyLinkedList struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	Val  int
	Next *Node
	Pre  *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	// index无效
	if this.size == 0 || index < 0 || index >= this.size {
		return -1
	}
	// index为头尾节点
	if index == 0 {
		return this.head.Val
	}
	if index == this.size-1 {
		return this.tail.Val
	}
	// 下面选择更快的遍历方向, 查看index在偏头部还是偏尾部
	cur := this.head
	count := 0
	for cur != nil {
		if count == index {
			break
		}
		count++
		cur = cur.Next
	}
	return cur.Val
}

//头结点
func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{Val: val}
	if this.head != nil {
		node.Next = this.head
		this.head.Pre = node
		this.head = node
	} else {
		this.head = node
		this.tail = this.head
	}
	this.size++
}

//添加尾结点
func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{Val: val}
	if this.tail != nil {
		node.Pre = this.tail
		this.tail.Next = node
		this.tail = node
	} else {
		this.tail = node
		this.head = this.tail
	}
	this.size++

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	//头结点
	if index > this.size {
		return
	}
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.size {
		this.AddAtTail(val)
		return
	}
	cur := this.head
	count := 0
	for cur != nil && count < index {
		count++
		cur = cur.Next
	}
	// 在当前节点cur的前面插入
	this.size++
	node := &Node{Val: val, Next: cur, Pre: cur.Pre}
	cur.Pre.Next = node
	cur.Pre = node

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	//链表为空
	// 如果索引无效 || 链表为空
	if this.size == 0 || index < 0 || index >= this.size {
		return
	}
	if index == 0 {
		// 如果删除的是头结点
		this.head = this.head.Next
	} else if index == this.size-1 {
		// 如果删除的是尾结点
		this.tail = this.tail.Pre
	} else {
		cur := this.head
		count := 0
		for cur != nil && count < index {
			count++
			cur = cur.Next
		}
		// 找到要删除的节点cur了
		cur.Next.Pre = cur.Pre
		cur.Pre.Next = cur.Next
	}
	this.size--

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
