package linkedlist

import "sort"

/*
输入一个链表，按链表从尾到头的顺序返回一个ArrayList。

遍历链表，取出所有值，然后再反转数组输出
*/

func printListFromTailToHead(head *ListNode) []int {
	// write code here
	tmp := []int{}
	res := []int{}
	for head != nil {
		tmp = append(tmp, head.Val)
		head = head.Next
	}
	l := len(tmp) - 1
	for l > -1 {
		res = append(res, tmp[l])
		l--
	}
	return res
}

/*
给定一个无序单链表，实现单链表的排序(按升序排序)
*/
func sortInList(head *ListNode) *ListNode {
	mid := head
	values := []int{}
	for mid != nil {
		values = append(values, mid.Val)
		mid = mid.Next
	}
	sort.Ints(values)
	mid = head
	for _, v := range values {
		mid.Val = v
		mid = mid.Next
	}
	return head
}
