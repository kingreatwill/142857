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

// 最优时间.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var root = &ListNode{}
	result := root
	var flag = 0 //记录进位
	for l1 != nil || l2 != nil || flag != 0 {
		var lival int
		if l1 != nil { //先判断链表是否为空
			lival = l1.Val
		} else {
			lival = 0
		}
		var l2val int
		if l2 != nil { //先判断链表是否为空
			l2val = l2.Val
		} else {
			l2val = 0
		}
		temp := &ListNode{Val: (lival + l2val + flag) % 10} //记录暂时数据
		flag = (lival + l2val + flag) / 10
		result.Next = temp
		result = result.Next
		if l1 != nil {
			l1 = l1.Next //下个节点
		}
		if l2 != nil {
			l2 = l2.Next //下个节点
		}
	}
	return root.Next
}

// 最优内存.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var root = &ListNode{}
	var prev = root
	isPlus := false
	var next *ListNode
	for l1 != nil {
		next = l1
		l1 = l1.Next
		next.Next = nil
		if isPlus {
			next.Val += 1
			isPlus = false
		}
		if l2 != nil {
			next.Val += l2.Val
			l2 = l2.Next
		}
		if next.Val >= 10 {
			next.Val %= 10
			isPlus = true
		}
		prev.Next = next
		prev = next
	}
	for l2 != nil {
		next = l2
		l2 = l2.Next
		next.Next = nil
		if isPlus {
			next.Val += 1
			isPlus = false
		}
		if next.Val >= 10 {
			next.Val %= 10
			isPlus = true
		}
		prev.Next = next
		prev = next
	}
	if isPlus {
		next = &ListNode{}
		next.Val = 1
		isPlus = false
		prev.Next = next
		prev = next
	}
	return root.Next
}
