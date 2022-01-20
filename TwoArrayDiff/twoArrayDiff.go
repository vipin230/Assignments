package main

import (
	"fmt"
	"strings"
)

func main() {
	var arr1, arr2 string
	fmt.Println("Enter two arrays to compare:")
	fmt.Scanln(&arr1)
	fmt.Scanln(&arr2)

	map1 := map[string]int{}

	str1 := strings.Split(arr1, "")
	str2 := strings.Split(arr2, "")

	for _, val := range str2 {
		if map1[val] == 0 {
			map1[val] = 1
		} else {
			map1[val]++
		}
	}

	result := []string{}
	for _, val := range str1 {
		if num, ok := map1[val]; ok && num > 0 {
			map1[val]--
		} else {
			result = append(result, val)
		}
	}

	fmt.Printf("Following numbers from first array(%s) not available in 2nd array(%s): %s", str1, str2, result)
}
