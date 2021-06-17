package tree

/*
操作给定的二叉树，将其变换为源二叉树的镜像。
比如：    源二叉树
            8
           /  \
          6   10
         / \  / \
        5  7 9 11
        镜像二叉树
            8
           /  \
          10   6
         / \  / \
        11 9 7  5

思路：
	从根节点遍历，将左右子节点放入队列，然后调换左右子节点，继续遍历队列
*/
func Mirror(pRoot *TreeNode) *TreeNode {
	if pRoot == nil {
		return pRoot
	}
	cur := []*TreeNode{pRoot}
	next := []*TreeNode{}
	for i := 0; len(cur) > 0; {
		node := cur[i]
		if node.Left != nil {
			next = append(next, node.Left)
		}
		if node.Right != nil {
			next = append(next, node.Right)
		}
		left := node.Left
		node.Left = node.Right
		node.Right = left
		if i == len(cur)-1 {
			cur = next
			next = []*TreeNode{}
			i = 0
			continue
		}
		i++
	}
	return pRoot
}
