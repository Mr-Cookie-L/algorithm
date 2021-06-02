package stack

// 用两个栈来实现一个队列，完成队列的Push和Pop操作。 队列中的元素为int类型。

var stack1 []int
var stack2 []int

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int {
	if len(stack2) == 0 {
		stack2 = append(stack2, stack1...)
		stack1 = make([]int, 0)
	}
	if len(stack2) > 0 {
		res := stack2[0]
		if len(stack2) >= 2 {
			stack2 = stack2[1:]
		} else {
			stack2 = make([]int, 0)
		}

		return res
	}
	return -1
}
