package main

import "fmt"

func fib(x int) int {
	if x <= 1 {
		return 1
	}
	return fib(x-1) + fib(x-2)
}

func main() {
	var x int
	fmt.Println("Input number for fibonacci:")
	fmt.Scanf("%d", &x)
	result := fib(x)
	fmt.Println("Fibonacci of", x, "is:", result)
}
