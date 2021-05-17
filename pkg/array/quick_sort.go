package array

// QuickSort 快排
func QuickSort(array []int, left, right int) {
	if !(left < right) {
		return
	}
	i, j := left, right
	key := array[(i+j)/2]
	for i <= j {
		for array[i] < key {
			i++
		}
		for array[j] > key {
			j--
		}
		if i <= j {
			array[i], array[j] = array[j], array[i]
			i++
			j--
		}
	}

	if left < j {
		QuickSort(array, left, j)
	}
	if right > i {
		QuickSort(array, i, right)
	}
}
