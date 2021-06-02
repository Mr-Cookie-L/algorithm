package linkedlist

/*
type ListNode struct {
	Val  int
	Next *ListNode
}
*/

// 判断一个链表是否存在环
/*
思路：快慢指针
	初始化两个指针，快指针走两步，慢指针走一步，如果存在环，两个指针一定会相遇
*/

func HasCircle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}

	return false
}

// 变种，同样使用快慢指针
// 给一个链表，若其中包含环，请找出该链表的环的入口结点，否则，输出null
/*
思路：
	同上题一样使用快慢指针先把相交节点获取到，
	此时，快指针走过的路程是慢指针的2倍，且他们相遇的地方一定在环内，
	此时，以环入口节点、相交节点为界，把路程分为
	A-环外长度
	B-从环入口到相交点长度
	C-从相交点到环入口剩余长度
	由于快指针路程是慢指针的两倍，所以有 2(A+B) = A+B+C+B
	即 A = C
	所以此时，将其中一个指针指向头，然后两个指针同步，再次相交点就是B点，即环入口
*/
func EntryNodeOfLoop(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	slow2 := head

	for slow2 != slow {
		slow = slow.Next
		slow2 = slow2.Next
	}
	return slow
}
