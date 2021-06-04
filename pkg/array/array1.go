package array

import (
	"sort"
	"strconv"
	"strings"
)

/*
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
输入一个非递减排序(后一个元素>=前一个元素)的数组的一个旋转，输出旋转数组的最小元素。
NOTE：给出的所有元素都大于0，若数组大小为0，请返回0。

思路：使用二分法，递归，逐步缩小数组的范围，最终找到最小的值（也就是旋转的原点）
*/
func minNumberInRotateArray(rotateArray []int) int {
	// write code here
	if len(rotateArray) == 0 {
		return 0
	}

	left, right := 0, len(rotateArray)-1

	for left < right {
		if rotateArray[left] < rotateArray[right] {
			return rotateArray[left]
		}
		mid := (left + right) / 2
		if rotateArray[mid] < rotateArray[right] {
			right = mid
		} else if rotateArray[mid] > rotateArray[right] {
			left = mid + 1
		} else {
			right--
		}
	}
	return rotateArray[left]
}

/*
在一个长度为n的数组里的所有数字都在0到n-1的范围内。 数组中某些数字是重复的，但不知道有几个数字是重复的。也不知道每个数字重复几次。请找出数组中任一一个重复的数字。 例如，如果输入长度为7的数组[2,3,1,0,2,5,3]，那么对应的输出是2或者3。存在不合法的输入的话输出-1
*/
func Duplicate(numbers []int) int {
	// 使用map来计数
	m := map[int]int{}
	for _, n := range numbers {
		m[n] += 1
		if m[n] > 1 {
			return n
		}
	}
	return -1
}

/*
给定一个数组A[0,1,...,n-1],请构建一个数组B[0,1,...,n-1],其中B中的元素B[i]=A[0]*A[1]*...*A[i-1]*A[i+1]*...*A[n-1]。不能使用除法。（注意：规定B[0] = A[1] * A[2] * ... * A[n-1]，B[n-1] = A[0] * A[1] * ... * A[n-2];）
对于A长度为1的情况，B无意义，故而无法构建，因此该情况不会存在。
*/
/*
分析：
	由于不能用除法，且B[i]=A[0]*A[1]*...*A[i-1]*A[i+1]*...*A[n-1]，我们可以把b[i]分成左右两部分，即
	B[i] = left[i] * right[i]
	left[i] = A[0]*A[1]*...*A[i-1]
	right[i] = A[i+1]*A[i+2]*...*A[n-1]
	故
	left[i+1] = A[0]*A[1]*...*A[i] = left[i]*A[i]
	right[i+1] = A[i+2]*A[i+3]*...*A[n-1] => right[i] = right[i+1]*A[i+1]
	所以，我们可以先计算left，再计算right
	B[0]	1		A[1]	A[2]	A[3]	...    A[n-1]
	B[1]	A[0]	1		A[2]	A[3]	...    A[n-1]
	B[2]	A[0]	A[1]	1		A[3]	...    A[n-1]
	B[3]	A[0]	A[1]	A[2]	1		...    A[n-1]
	B[n-1]	A[0]	A[1]	A[2]	A[3]	...    1
*/
func multiply(A []int) []int {
	l := len(A)
	B := make([]int, l)
	tmp := 1
	for i := 0; i < l; i++ {
		B[i] = tmp
		tmp *= A[i]
	}
	tmp = 1
	for i := l - 1; i >= 0; i-- {
		B[i] *= tmp
		tmp *= A[i]
	}
	return B
}

/*
输入一个数组，长度为100，其中值的范围都是0-10，求该数组的99分位数（array[有序数组长度*0.99]）

思路：由于需要数组有序，且值范围都在0-10之间，使用map统计其出现次数，然后找到对应的索引的值即可
*/
func Index99(input []int) int {
	index99 := int(float64(len(input)) * 0.99)
	m := map[int]int{}
	for _, n := range input {
		m[n]++
	}
	j := 0
	for i := 0; i <= 10; i++ {
		j += m[i]
		if j >= index99 {
			return i
		}
	}
	return 0
}

/*
输入一个无序的数组，将数组排成最小的字符串
例如输入[1,21,2]，输出 1212

思路：重点在于字符串状态下 a+b < b+a  那么a就排在前面

go语言可以直接使用字符串进行比较，比较时会将string转化为[]byte，然后依次比较每个元素
例如 "101" < "1001" 返回是 false，因为string的第三位转化为[]byte时，1 > 0，因此就会返回false
*/
func PrintMinNumber(numbers []int) string {
	strList := []string{}
	for _, n := range numbers {
		strList = append(strList, strconv.Itoa(n))
	}
	sort.Slice(strList, func(i, j int) bool {
		return strList[i]+strList[j] < strList[j]+strList[i]
	})
	return strings.Join(strList, "")
}

/*
给定一个数组arr，返回子数组的最大累加和
例如，arr = [1, -2, 3, 5, -2, 6, -1]，所有子数组中，[3, 5, -2, 6]可以累加出最大的和12，所以返回12.
题目保证没有全为负数的数据
[要求]
时间复杂度为O(n)O(n)，空间复杂度为O(1)O(1)

思路：
	设置两个值，一个代表全局最大和(max)，一个代表当前和(cur)
	如果当前和 < 0，则将其=0
	如果当前和 > max，则max = cur
	如此循环完数组，max一定是最大和
*/
func MaxsumofSubarray(arr []int) int {
	max := 0
	cur := 0
	if len(arr) == 1 {
		return arr[0]
	}
	for _, n := range arr {
		cur += n
		if cur < 0 {
			cur = 0
		} else if cur > max {
			max = cur
		}
	}
	return max
}

/*
给出两个有序的整数数组A和B ，请将数组B合并到数组A中，变成一个有序的数组
注意：
可以假设A数组有足够的空间存放B数组的元素，A和B中初始的元素数目分别为m和n

思路：
	由于不会返回新的数组，所以要在A数组上直接操作
	合并操作，如果考虑到时间复杂度，肯定要用归并排序，不能使用双层循环
	如果从头开始遍历，那A数组的元素向后移动肯定很困难
	所以可以考虑从尾遍历，既然A有足够空间，那可以从m+n的位置入手去修改数据
*/
func Merge(A []int, m int, B []int, n int) {
	a, b, i := m-1, n-1, m+n-1
	for i >= 0 {
		if a >= 0 && b >= 0 {
			if A[a] > B[b] {
				A[i] = A[a]
				a--
			} else {
				A[i] = B[b]
				b--
			}
		} else if a < 0 && b >= 0 {
			A[i] = B[b]
			b--
		} else {
			break
		}
		i--
	}
}

/*
上题的简单版本，A、B为有序数组(升序)，输出合并后的有序数组

思路：
	归并排序思路，循环并判断值的大小
*/
func MergeAndReturn(A []int, B []int) []int {
	i, j := 0, 0
	res := []int{}
	for i < len(A)-1 && j < len(B)-1 {
		if A[i] > B[j] {
			res = append(res, B[j])
			j++
		} else {
			res = append(res, A[i])
			i++
		}
	}
	if i >= len(A) {
		res = append(res, B[j:]...)
	}
	if j >= len(B) {
		res = append(res, A[i:]...)
	}
	return res
}
