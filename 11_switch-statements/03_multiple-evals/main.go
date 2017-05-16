package main

import "fmt"

func main() {
	switch "Tim" {
	case "Tim", "Jimmy":
		fmt.Println("Sup Tim... er or Jimmy")
	case "Marcus":
		fmt.Println("Sup Marcus")
	case "Luke":
		fmt.Println("Sup Luke")
	default:
		fmt.Println("You have no friends")
	}
}