package main

import (
	"fmt"

	"github.com/Mr-Cookie-L/algorithm/pkg/array"
)

func main() {
	a := []int{1, 2, 3, 4, 2, 2, 2, 2}
	b, c := array.Maxelement(a)
	fmt.Println("b,c------", b, c)
}
