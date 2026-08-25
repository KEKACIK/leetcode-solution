package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	var old *ListNode
	cur := slow
	for cur != nil {
		next := cur.Next
		cur.Next = old
		old = cur
		cur = next
	}

	lHead := head
	lNew := old
	for lHead != nil && lNew != nil {
		fmt.Println(lHead.Val, lNew.Val)
		if lHead.Val != lNew.Val {
			return false
		}
		lHead = lHead.Next
		lNew = lNew.Next
	}

	return true
}
