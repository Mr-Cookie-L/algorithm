package main

import (
	"fmt"

	"github.com/Mr-Cookie-L/algorithm/pkg/array"
)

func main() {
	a := []int{5, 2, 3, 1, 4}
	array.QuickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
