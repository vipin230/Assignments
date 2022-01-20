package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{10, 5, 10}
	//arr := []int{10, 10, 10}
	//arr := []int{12, 35, 1, 10, 34, 1}

	fmt.Println("Array:", arr)
	sort.Ints(arr)
	i := len(arr) - 1
	for i > 0 {
		if arr[i-1] != arr[i] {
			fmt.Println("Second Highest value in array is:", arr[i-1])
			return
		}
		i--
	}
	fmt.Println("There is no second largest element")
}
