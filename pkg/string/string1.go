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
