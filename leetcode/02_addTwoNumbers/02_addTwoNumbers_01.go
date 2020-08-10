package main

import "fmt"

func main() {
	r := addTwoNumbers(&ListNode{Val: 2, Next: &ListNode{4, &ListNode{3, nil}}}, &ListNode{Val: 5, Next: &ListNode{6, &ListNode{4, nil}}})
	fmt.Printf("%v\n", r)
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	curr := result
	carry := 0
	first := true
	for l1 != nil || l2 != nil || carry > 0 {
		if !first {
			curr.Next = new(ListNode)
			curr = curr.Next
		}
		first = false
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		curr.Val = carry % 10
		carry /= 10
	}
	return result
}
