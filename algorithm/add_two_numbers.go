package algorithm

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
https://leetcode.com/problems/add-two-numbers/
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
 */
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, next *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum = sum + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum = sum + l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		if head == nil {
			head = &ListNode{sum % 10, nil}
			next = head
		} else if next != nil {
			next.Next = &ListNode{sum % 10, nil}
			next = next.Next
		}

	}
	if carry != 0 {
		next.Next = &ListNode{carry, nil}
	}
	return head
}

func BuildLink(input []int) *ListNode {
	var head, next *ListNode
	for _, v := range input {
		if head == nil {
			head = &ListNode{v, nil}
			next = head
			continue
		}
		if next.Next == nil {
			next.Next = &ListNode{v, nil}
			next = next.Next
		}
	}
	return head
}
