package main

import "fmt"

func main() {
	if false {
		fmt.Println("This will not print")
	} else {
		fmt.Println("This will print though")
	}
}