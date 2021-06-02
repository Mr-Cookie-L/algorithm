package array

// QuickSort 快排
func QuickSort(array []int, left, right int) {
	if !(left < right) {
		return
	}
	// 两个指针
	i, j := left, right
	// 找到中间值
	key := array[(i+j)/2]
	for i <= j {
		// 找到左边比key大的值
		for array[i] < key {
			i++
		}
		// 找到右边比key小的值
		for array[j] > key {
			j--
		}
		// 判断一下两个指针没有交错而过或者相等
		if i <= j {
			// 交换左右两边的数字
			array[i], array[j] = array[j], array[i]
			i++
			j--
		}
	}

	// 将左数组递归进入循环继续排序
	if left < j {
		QuickSort(array, left, j)
	}
	// 将右数组递归进入循环继续排序
	if right > i {
		QuickSort(array, i, right)
	}
}
