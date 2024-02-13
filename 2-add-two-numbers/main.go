package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Create the first list: 6 -> 0 -> 1
	l1 := &ListNode{Val: 6}
	l1.Next = &ListNode{Val: 0}
	l1.Next.Next = &ListNode{Val: 1}

	// Create the second list: 5 -> 1 -> 9
	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 1}
	l2.Next.Next = &ListNode{Val: 9}

	// Call the addTwoNumbers function and print the result
	result := addTwoNumbers(l1, l2)
	for result != nil {
		fmt.Printf("%d -> ", result.Val)
		result = result.Next
	}

}

// addTwoNumbers takes two singly-linked lists representing non-negative integers
// and returns their sum as a singly-linked list.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Initialize variables to track the carry and the head of the result list
	var carry int
	var head = &ListNode{}
	var resultTail = head

	// Iterate through the lists until both are exhausted and there's no carry left
	for l1 != nil || l2 != nil || carry != 0 {
		// Get the current values from the 1st list
		v1 := 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		// Get the current values from the 2nd list
		v2 := 0
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		// Calculate the sum and update the carry
		sum := v1 + v2 + carry
		carry = sum / 10

		// Create a new node for the sum and append it to the result list
		newNode := &ListNode{Val: sum % 10}
		resultTail.Next = newNode
		resultTail = newNode
	}

	// Return the head of the result list (excluding the dummy head node)
	return head.Next
}
