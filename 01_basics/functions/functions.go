package main

import (
	"fmt"
)

func addLongHand(x int, y int) int {
	return x + y
}

func addShorthand(x, y int) int {
	return x + y
}
func main() {
	fmt.Println(addLongHand(42, 13))
	fmt.Println(addShorthand(42, 13))
}
