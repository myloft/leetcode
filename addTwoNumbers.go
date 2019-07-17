/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := l2
	sum := 0
	carry := 0
	for {
		sum = l1.Val + l2.Val + carry
		l2.Val = sum % 10
		carry = sum / 10
		if l1.Next == nil && l2.Next == nil && carry == 0 {
			break
		} else {
			if l1.Next == nil {
				l1.Next = &ListNode{0, nil}
			}
			if l2.Next == nil {
				l2.Next = &ListNode{0, nil}
			}
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return res
}