package main

import "fmt"

func main() {
	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("An error occured")
	} else {
		fmt.Printf("Hello, %v!!!\n", name)
	}
}