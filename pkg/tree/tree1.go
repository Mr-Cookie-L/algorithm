package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
从上往下打印出二叉树的每个节点，同层节点从左至右打印。

解题思路：
	按照层去遍历，顺序遍历，首先想到使用队列来实现，将root的左右子节点分别放在队列中
	然后循环队列，将下一层的子节点也放入队列中，将左右子树拿到的值加入到数组中，一次类推
	利用了队列的先进先出的特点
*/
func PrintFromTopToBottom(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	cur := []*TreeNode{root}
	next := []*TreeNode{}
	for i := 0; len(cur) > 0; {
		node := cur[i]
		res = append(res, node.Val)
		if node.Left != nil {
			next = append(next, node.Left)
		}
		if node.Right != nil {
			next = append(next, node.Right)
		}
		if i == len(cur)-1 {
			cur = next
			next = []*TreeNode{}
			i = 0
			continue
		}
		i++
	}
	return res
}
