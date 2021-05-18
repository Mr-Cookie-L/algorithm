package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
反转链表
思路：
	需要一个中间值，记录当前节点的NEXT节点，
	然后把当前节点的NEXT指针指向上一个节点，
	然后再把当前节点移动到NEXT上，以此类推
*/
func ReverseList(curr *ListNode) *ListNode {
	// write code here
	var prev *ListNode
	next := new(ListNode)

	for {
		if curr == nil {
			break
		}
		// 1.先记录下一个节点
		next = curr.Next
		// 2.将当前节点的next反转
		curr.Next = prev
		// 3.反转后的上一节点前移一格
		prev = curr
		// 4.将当前节点前移一格
		curr = next
	}
	return prev
}
