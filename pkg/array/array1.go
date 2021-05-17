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
