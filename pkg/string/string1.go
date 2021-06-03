package string

/*
在一个字符串(0<=字符串长度<=10000，全部由字母组成)中找到第一个只出现一次的字符,
并返回它的位置, 如果没有则返回 -1（需要区分大小写）.（从0开始计数）
*/
// 直接使用map，或者使用数组代替map也可以
// 这里设计到rune类型，建议了解一下

func FirstNotRepeatingChar(str string) int {
	hashTime := [256]int{0}
	for _, s := range str {
		hashTime[s]++
	}
	for i, s := range str {
		if hashTime[s] == 1 {
			return i
		}
	}
	return -1
}

/*
找出一个字符串符合条件的最长子串
该子串必须有完整的小括号包裹，且子串内没有小括号
例如：
输入  (123)(333)(()123123)(1111)
输出  1111

思路：
目前还没有找到相关题目的解法，只能使用自己思考出的方法
	1. 遇到 ( 时，则标记为子串的开头，将之前记录的子串清空
	2. 遇到 ) 时，判断子串开头是否为 ( ，如果是，则判断当前子串与上一个子串的长度，长则替换，短则清空开头和记录的子串；如果不是则不需要记录
	3. 直到循环完成，得出的最终子串便为最长子串
*/
func LongestSub(input string) string {
	res := ""
	start := ""
	mid := ""
	for i := 0; i < len(input); i++ {
		s := string(input[i])
		switch s {
		case "(":
			start = "("
			mid = ""
		case ")":
			if start != "(" {
				continue
			}
			if len(mid) > len(res) {
				res = mid
			} else {
				mid = ""
				start = ""
			}
		default:
			if start == "(" {
				mid += s
			}
		}
	}
	return res
}

/*
输入一个只有()的字符串，找出最长的合法子串的长度，即每一个 ( 都能对应一个 )
例如
输入 ())(())
输出 4
*/
func LongestValidParentheses(s string) int {
	res := 0
	stack := make([]int, len(s)+1)
	// 边界问题
	stack[0] = -1
	top := 0
	for i, b := range s {
		str := string(b)
		if str == "(" {
			top++
			stack[top] = i
		} else {
			top--
			if top == -1 {
				top = 0
				stack[top] = i
			} else {
				if res < i-stack[top] {
					res = i - stack[top]
				}
			}
		}
	}
	return res
}
