package main

import (
	"fmt"

	"github.com/Mr-Cookie-L/algorithm/pkg/array"
)

func main() {
	a := []int{1, 3, 4, 2, 6, 8, 2, 0, 19, 9}
	array.QuickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
