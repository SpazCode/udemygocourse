package main

import (
	"fmt"
)

func main() {
	a := 43
	
	fmt.Println(a)
	fmt.Println(&a)

	// b is a int pointer
	var b *int = &a

	// b is the same address as a
	fmt.Println(b)
	// To see the value in the pointer use '*'
	fmt.Println(*b)

	// reassign the value in the address of b
	*b = 42
	fmt.Println(a)
}