package main

import (
	"fmt"

	"github.com/Mr-Cookie-L/algorithm/pkg/queue"
)

func main() {
	a := []int{3, 2, 4, 6, 1, 5, 8, 7, 2}
	b := queue.GetLeastNumbers(a, 3)
	fmt.Println("b------", b)
}
