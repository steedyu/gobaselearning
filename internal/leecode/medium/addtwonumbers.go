package medium


/*
https://leetcode.com/problems/add-two-numbers/description/

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
 */

type ListNode struct {
	Val int
	Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var retlistnode *ListNode = new(ListNode)
	tmp := retlistnode

	v1 := l1.Val
	v2 := l2.Val
	m := 0

	for {
		v := v1 + v2 + m

		m = 0
		v1 = 0
		v2 = 0

		if v > 9 {
			m = v / 10
			v = v % 10
		}
		tmp.Val = v

		if l1.Next == nil && l2.Next == nil && m == 0 {
			break
		}else {
			if l1.Next != nil {
				l1 = l1.Next
				v1 = l1.Val
			}
			if l2.Next != nil {
				l2 = l2.Next
				v2 = l2.Val
			}
		}

		tmp.Next = new(ListNode)
		tmp = tmp.Next
	}

	return retlistnode

}
