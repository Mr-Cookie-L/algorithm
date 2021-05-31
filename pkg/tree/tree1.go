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
	同样也可以使用两个数组，进行循环
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

/*
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历的结果。如果是则返回true,否则返回false。
假设输入的数组的任意两个数字都互不相同。（ps：我们约定空树不是二叉搜素树）

输入：
[4,8,6,12,16,14,10]
返回值：
true

解题思路：
	由于是后序遍历，所以每个子树的根节点都在子数组的最后一位，因此可以使用数组分割的方式，
	判断 数组左边所有的值 < 根节点值 < 数组右边所有的值，且进行递归
*/
func VerifySquenceOfBST(sequence []int) bool {
	l := len(sequence)
	if l < 1 {
		return false
	}
	// 获取根节点
	root := sequence[l-1]
	// 定义左子数组
	var left []int
	i := 0
	// 找到左子树
	for sequence[i] < root && i < l-1 {
		left = append(left, sequence[i])
		i++
	}
	// 验证右子树是否符合条件
	for _, s := range sequence[i : l-1] {
		if s <= root {
			return false
		}
	}

	switch {
	case len(left) > 0 && len(sequence[i:l-1]) > 0:
		return VerifySquenceOfBST(left) && VerifySquenceOfBST(sequence[i:l-1])
	case len(left) == 0 && len(sequence[i:l-1]) > 0:
		return VerifySquenceOfBST(sequence[i : l-1])
	case len(left) > 0 && len(sequence[i:l-1]) == 0:
		return VerifySquenceOfBST(left)
	default:
		return true
	}
}
