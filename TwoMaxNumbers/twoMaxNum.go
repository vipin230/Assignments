package main

import "fmt"

func main() {
	//arr := []int{10, 5, 10}
	arr := []int{5, 34, 78, 2, 45, 1, 99, 23}
	//arr := []int{10, 10, 10}
	//arr := []int{12, 35, 1, 10, 34, 1}

	highest := 0
	second := 0

	for _, val := range arr {
		if highest < val {
			second = highest
			highest = val
		} else if second < val && val != highest {
			second = val
		}
	}

	fmt.Println("Highest value in array:", highest)
	if second != 0 {
		fmt.Println("Second highest value in array:", second)
	} else {
		fmt.Println("There is no second highest number in an integer array.")
	}

}
