package main

import (
	"fmt"

	"github.com/Mr-Cookie-L/algorithm/pkg/linkedlist"
)

func main() {
	link1 := &linkedlist.ListNode{
		Val: 1,
		Next: &linkedlist.ListNode{
			Val: 4,
			Next: &linkedlist.ListNode{
				Val: 5,
			},
		},
	}
	link2 := &linkedlist.ListNode{
		Val: 2,
		Next: &linkedlist.ListNode{
			Val: 3,
			Next: &linkedlist.ListNode{
				Val: 6,
			},
		},
	}
	res := linkedlist.Merge(link1, link2)
	for {
		if res.Next == nil {
			break
		}
		fmt.Println("res.val-------------", res.Val)
		res = res.Next
	}
}
