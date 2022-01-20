package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var input string
	fmt.Println("Enter the input:")
	fmt.Scan(&input)

	var postfix string
	var stack []string

	for _, val := range strings.Split(input, "") {
		if isAlphaNum(val) {
			postfix += val
		} else if val == "(" {
			stack = append(stack, val)
		} else if val == "^" {
			stack = append(stack, val)
		} else if val == ")" {
			for len(stack) != 0 && stack[len(stack)-1] != "(" {
				postfix += stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		} else {
			if prec(val) > prec(stack[len(stack)-1]) {
				stack = append(stack, val)
			} else {
				for len(stack) != 0 && prec(val) <= prec(stack[len(stack)-1]) {
					postfix += stack[len(stack)-1]
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, val)
			}
		}
	}

	for len(stack) != 0 {
		postfix += stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}

	fmt.Println("Infix:", input, "\nPostfix:", postfix)

}

func prec(s string) int {
	if s == "^" {
		return 3
	} else if (s == "/") || (s == "*") {
		return 2
	} else if (s == "+") || (s == "-") {
		return 1
	} else {
		return -1
	}
}

func isAlphaNum(x string) bool {
	result, _ := regexp.MatchString(`[a-zA-Z0-9]`, x)
	return result
}
