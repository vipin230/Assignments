package main

import (
	"fmt"
)

func main() {
	var binary int

	fmt.Println("Input Binary:")
	fmt.Scanf("%d", &binary)
	binary1 := binary
	decimal, rem, pow := 0, 0, 1
	for binary != 0 {
		rem = binary % 10
		decimal = decimal + rem*pow
		binary = binary / 10
		pow = pow * 2
	}

	fmt.Println("Binary:", binary1, "\nDecimal:", decimal)
}
