package main

import (
	"fmt"

	"github.com/Mr-Cookie-L/algorithm/pkg/tree"
)

func main() {
	// {10,6,14,4,8,12,16}
	root := &tree.TreeNode{
		Val: 10,
		Left: &tree.TreeNode{
			Val: 6,
			Left: &tree.TreeNode{
				Val: 4,
			},
			Right: &tree.TreeNode{
				Val: 8,
			},
		},
		Right: &tree.TreeNode{
			Val: 14,
			Left: &tree.TreeNode{
				Val: 12,
			},
			Right: &tree.TreeNode{
				Val: 16,
			},
		},
	}
	a := tree.PrintFromTopToBottom(root)
	fmt.Println(a)
}
