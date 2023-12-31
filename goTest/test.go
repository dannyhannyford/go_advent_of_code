package main

import (
	"fmt"
)

func main() {
	a := 10
	b := &a
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("*b:", *b)

	*b = 20
	fmt.Println("a after update:", a)
}