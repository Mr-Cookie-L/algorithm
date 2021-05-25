package linkedlist

// 输入两个单调递增的链表，输出两个链表合成后的链表，当然我们需要合成后的链表满足单调不减规则。
/*
1. 从头部开始遍历两个链表
2. 对比两个链表元素的大小
3. 小的加入到新链表，且指针往前进一位，以此类推，直到有一个链表的next=nil
*/

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	res := &ListNode{}
	mid := res
	for {
		if pHead1 == nil {
			mid.Next = pHead2
			break
		}
		if pHead2 == nil {
			mid.Next = pHead1
			break
		}
		if pHead1.Val >= pHead2.Val {
			next := pHead2.Next
			pHead2.Next = mid.Next
			mid.Next = pHead2
			pHead2 = next
		} else {
			next := pHead1.Next
			pHead1.Next = mid.Next
			mid.Next = pHead1
			pHead1 = next
		}
		mid = mid.Next
	}
	return res.Next
}
