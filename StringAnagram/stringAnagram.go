package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1, str2 string
	fmt.Println("Input 2 strings to check Anagram:")
	fmt.Scanf("%s %s", &str1, &str2)

	if isAnagram(str1, str2) {
		fmt.Println("String 1:", str1, "& String 2:", str2, "\nBoth strings are Anagrams")
	} else {
		fmt.Println("String 1:", str1, "& String 2:", str2, "\nBoth strings are NOT Anagrams")
	}

}

func isAnagram(str1, str2 string) bool {

	if len(str1) != len(str2) {
		return false
	}

	map1 := make(map[string]int, 0)

	for _, val := range strings.Split(str1, "") {
		if map1[val] == 0 {
			map1[val] = 1
		} else {
			map1[val]++
		}
	}

	for _, val := range strings.Split(str2, "") {
		if v, ok := map1[val]; ok && v > 0 {
			map1[val]--
		} else {
			return false
		}
	}

	return true
}
