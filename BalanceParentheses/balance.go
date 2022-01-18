package main

import (
	"fmt"
)

func main() {
	var params string
	fmt.Println("Input:")
	fmt.Scanf("%s", &params)

	fmt.Println(params, "-", checkParams(params))

}

func checkParams(params string) string {
	paramsMap := map[rune]string{'}': "{", ')': "(", ']': "["}

	stack := []string{}

	for _, val := range params {
		if val == '}' || val == ')' || val == ']' {
			if len(stack) >= 1 && paramsMap[val] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return "Not Balanced"

			}
		} else {
			stack = append(stack, string(val))
		}
	}
	if len(stack) == 0 {
		return "Balanced"
	} else {
		return "Not Balanced"
	}
}
