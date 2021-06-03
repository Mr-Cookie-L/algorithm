package main

import (
	"fmt"

	"github.com/Mr-Cookie-L/algorithm/pkg/string"
)

func main() {
	a := "(()(())"
	b := string.LongestValidParentheses(a)
	fmt.Println("b-----------", b)
}
