package queue

/*
给定一个数组，找出其中最小的K个数。例如数组元素是4,5,1,6,2,7,3,8这8个数字，则最小的4个数字是1,2,3,4。如果K>数组的长度，那么返回一个空的数组

输入：
[4,5,1,6,2,7,3,8,1],4
返回值：
[1,2,3,4]

思路：
	1. 可以使用直接排序的手段，为数组排序，然后直接返回
	// TODO 这个方法不太对，应该用单调队列或者栈，但是时间复杂度会增加
	2. 使用队列
		维护一个长度为k的队列，如果长度小于快k，则直接入队
		如果队满，则跟队首元素比较，如果小于，则队首元素出队，元素入队
*/
func GetLeastNumbers(arr []int, k int) []int {
	if len(arr) < k {
		return []int{}
	}
	queue := []int{}
	for _, i := range arr {
		if len(queue) < k {
			queue = append(queue, i)
			continue
		}
		if queue[0] < i {
			continue
		}
		queue = queue[1:]
		queue = append(queue, i)
	}
	return queue
}
