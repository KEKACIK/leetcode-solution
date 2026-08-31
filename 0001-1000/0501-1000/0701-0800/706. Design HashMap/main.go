package main

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Node struct {
	Key   int
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Get(key int) int {
	curr := ll.Head
	for curr != nil {
		if curr.Key == key {
			return curr.Value
		}
		curr = curr.Next
	}

	return -1
}

func (ll *LinkedList) Put(key, value int) {
	curr := ll.Head
	for curr != nil {
		if curr.Key == key {
			curr.Value = value
			return
		}
		curr = curr.Next
	}

	newNode := &Node{
		Key:   key,
		Value: value,
		Next:  ll.Head,
	}
	ll.Head = newNode
}

func (ll *LinkedList) Remove(key int) {
	if ll.Head == nil {
		return
	}

	if ll.Head.Key == key {
		ll.Head = ll.Head.Next
		return
	}

	curr := ll.Head
	for curr.Next != nil {
		if curr.Next.Key == key {
			curr.Next = curr.Next.Next
			return
		}
		curr = curr.Next
	}
}

type MyHashMap struct {
	n       int
	bukkets []*LinkedList
}

func Constructor() MyHashMap {
	n := 991
	bukkets := make([]*LinkedList, n)

	for i := range bukkets {
		bukkets[i] = &LinkedList{}
	}

	return MyHashMap{
		n:       n,
		bukkets: bukkets,
	}
}

func (this *MyHashMap) hash(x int) int {
	return abs(x) % this.n
}

func (this *MyHashMap) Put(key int, value int) {
	n := this.hash(key)
	this.bukkets[n].Put(key, value)
}

func (this *MyHashMap) Get(key int) int {
	n := this.hash(key)
	return this.bukkets[n].Get(key)
}

func (this *MyHashMap) Remove(key int) {
	n := this.hash(key)
	this.bukkets[n].Remove(key)
}
