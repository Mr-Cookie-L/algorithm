package main

import (
	"fmt"
)

func main() {
	str := "abcdef"
	for _, s := range str {
		fmt.Println("s----------", string(s))
	}
	fmt.Println("str[1]-------------", str[0], str[1], str[2])
}
