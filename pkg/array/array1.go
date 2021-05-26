package array

/*
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
输入一个非递减排序(后一个元素>=前一个元素)的数组的一个旋转，输出旋转数组的最小元素。
NOTE：给出的所有元素都大于0，若数组大小为0，请返回0。
*/

/**
 *
 * @param rotateArray int整型一维数组
 * @return int整型
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
