package main

type Node struct {
	Value int
	Next  *Node
}

type MyLinkedList struct {
	Head *Node
	Size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Size {
		return -1
	}

	cur := this.Head
	for i := 1; i < index; i++ {
		cur = cur.Next
	}

	return cur.Value
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.Size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Size {
		return
	}

	this.Size++
	if index == 0 {
		newNode := &Node{
			Value: val,
			Next:  this.Head,
		}
		this.Head = newNode
		return
	}

	cur := this.Head
	for i := 0; i < index-1; i++ {
		cur = cur.Next
	}

	oldNext := cur.Next
	cur.Next = &Node{
		Value: val,
		Next:  oldNext,
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}

	this.Size--
	if index == 0 {
		this.Head = this.Head.Next
		return
	}

	cur := this.Head
	for i := 0; i < index-1; i++ {
		cur = cur.Next
	}

	cur.Next = cur.Next.Next
}
