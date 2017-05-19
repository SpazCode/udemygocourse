package main

import "fmt"

const (
	A = iota
	B = iota
	C = iota
)

const (
	D = iota
	E = iota
	F = iota
)

func main() {
	fmt.Println(A)
	fmt.Println(D)
	fmt.Println(B)
	fmt.Println(E)
	fmt.Println(C)
	fmt.Println(F)
}