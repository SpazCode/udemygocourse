package main

import "fmt"

func main() {
	switch "Medhi" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Medhi":
		fmt.Println("Wassup Medhi")
	case "Mike":
		fmt.Println("Wassup Mike")
	default:
		fmt.Println("You have no friends")
	}
	
}

/*
	no default fall through so no break needed
	fall through is optional
*/