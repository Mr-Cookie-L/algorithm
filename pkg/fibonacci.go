package pkg

/* 斐波那契数列又称为黄金分割数列, 由于是以兔子的繁殖为例子引入的，因此也叫“兔子数列”。
它指的是这样一个数列：0,1,1,2,3,5,8,13......从这组数可以很明显看出这样一个规律：
从第三个数开始，后边一个数一定是在其之前两个数的和。
*/

// 现在要求输入一个整数n，请你输出斐波那契数列的第n项（从0开始，第0项为0，第1项是1）。

// Fibonacci 递归版本
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// Fibonacci2 非递归版本
func Fibonacci2(n int) int {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		b = a + b
		a = b - a
	}
	return a
}
