package main

import "fmt"

func main() {
	switch "Marcus" {
	case "Tim":
		fmt.Println("Sup Tim")
	case "Marcus":
		fmt.Println("Sup Marcus")
		fallthrough
	case "Mehdi":
		fmt.Println("Sup Mehdi")
		fallthrough
	case "Johnny":
		fmt.Println("Sup Johnny")
	default:
		fmt.Println("You have no frineds")
	}
	
}