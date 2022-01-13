package main

import (
	"fmt"
)

func main() {
	// x := "22"
	var num int

	fmt.Println("Enter a number to check Palindrome:")
	fmt.Scanf("%d", &num)

	if isPalindrome(num) {
		fmt.Println(num, "is a Palindrome!!!")
	} else {
		fmt.Println(num, "is not a Palindrome")
	}
}

func isPalindrome(num int) bool {
	if num < 0 || (num > 0 && num%10 == 0) {
		return false
	}

	tmp := 0
	for num > tmp {
		tmp = tmp*10 + num%10
		num = num / 10
	}

	if num == tmp || num == tmp/10 {
		return true
	} else {
		return false
	}
}
