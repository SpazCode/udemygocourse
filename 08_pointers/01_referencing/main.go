package main

import "fmt"

func main() {
	a :=  43

	fmt.Println(a)
	fmt.Println(&a)

	// b is of type 'int pointer'
	// the pointer is the address of the variable
	// '*' defines that it is a pointer
	var b *int = &a

	fmt.Println(b)
}