package array

/*
需要从k个数的array中找到最大的两个数 分别用循环算法和递归算法来写
如，输入[1,1,2,2]
输出[2,1]
*/
// 循环
func MaxelementFor(array []int) (int, int) {
	var max, second int
	if len(array) < 1 {
		return max, second
	}
	for _, i := range array {
		if i > max {
			second = max
			max = i
		} else if i > second && i < max {
			second = i
		}
	}

	return max, second
}

// 递归
func Maxelement(array []int) (int, int) {
	if len(array) < 1 {
		return 0, 0
	}
	if len(array) == 1 {
		return array[0], 0
	}
	mid := len(array) / 2
	max1, sencond1 := Maxelement(array[0:mid])
	max2, sencond2 := Maxelement(array[mid:])
	if max1 > max2 {
		if sencond1 > max2 {
			return max1, sencond1
		}
		return max1, max2
	}
	if sencond2 > max1 {
		return max2, sencond2
	}
	return max2, max1
}
